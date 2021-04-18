package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/zeebo/errs"
)

// BoardStateKey formats a board state key for boardstate to user mapping.
func BoardStateKey(gameID, username string) string {
	return fmt.Sprintf("%s:%s", gameID, username)
}

// TODO: Need to determine how we format boardstate keys
// TODO: Need to switch Directory over to Redis storage

func (s *graphQLServer) BoardstatePosted(ctx context.Context, bs InputBoardState) (<-chan *BoardState, error) {
	_, ok := s.Directory[bs.GameID]
	if !ok {
		return nil, fmt.Errorf("failed to find game %s", bs.GameID)
	}

	board, err := boardStateFromInput(bs)
	if err != nil {
		return nil, fmt.Errorf("failed to get boardstate from input: %s", err)
	}
	log.Printf("boardstate from input: %+v", board)

	boardstates := make(chan *BoardState, 1)
	s.mutex.Lock()
	s.boardChannels[board.User.ID] = boardstates
	s.mutex.Unlock()
	boardstates <- board
	go func() {
		<-ctx.Done()
		s.mutex.Lock()
		delete(s.boardChannels, board.User.ID)
		s.mutex.Unlock()
	}()

	return boardstates, nil
}

func (s *graphQLServer) UpdateBoardState(ctx context.Context, input InputBoardState) (*BoardState, error) {
	log.Printf("UpdateBoardState#input: %+v", input)
	_, ok := s.Directory[input.GameID]
	if !ok {
		return nil, fmt.Errorf("failed to updated boardstate: game does not exist: %s", input.GameID)
	}

	bs, err := boardStateFromInput(input)
	if err != nil {
		return nil, errs.Wrap(err)
	}

	if err := s.Set(BoardStateKey(input.GameID, input.User.Username), bs); err != nil {
		return nil, fmt.Errorf("failed to persist boardstate: %s", err)
	}

	s.BoardstatePosted(ctx, input)

	return bs, nil
}

// Boardstates queries Redis for different boardstates per player or game
func (s *graphQLServer) Boardstates(ctx context.Context, gameID string, username *string) ([]*BoardState, error) {
	game, ok := s.Directory[gameID]
	if game == nil {
		log.Printf("game is nil: %s - %+v", gameID, s.Directory)
		return nil, errs.New("game does not exist")
	}
	if !ok {
		log.Printf("game %s does not exist: %+v", gameID, s.Directory)
		return nil, errs.New("game does not exist")
	}

	// if username is not provided, send all
	if username == nil {
		boardstates := []*BoardState{}
		for _, p := range game.PlayerIDs {
			board := &BoardState{}
			boardKey := BoardStateKey(game.ID, p.Username)
			err := s.Get(boardKey, &board)
			if err != nil {
				log.Printf("error fetching user boardstate from redis: %s", err)
			}
			boardstates = append(boardstates, board)
		}
		return boardstates, nil
	}

	boardstates := []*BoardState{}
	for _, p := range game.PlayerIDs {
		if p.Username == *username {
			board := &BoardState{}
			boardKey := BoardStateKey(game.ID, p.Username)
			err := s.Get(boardKey, &board)
			if err != nil {
				log.Printf("error fetching user boardstate from redis: %s", err)
			}

			boardstates = append(boardstates, board)
		}
	}

	if len(boardstates) == 0 {
		return []*BoardState{}, errs.New("no boardstate for user %s found", *username)
	}

	return boardstates, nil
}

func boardStateFromInput(bs InputBoardState) (*BoardState, error) {
	data, err := json.Marshal(bs)
	if err != nil {
		return nil, errs.New("failed to marshal input game: %s", err)
	}
	new := &BoardState{}
	err = json.Unmarshal(data, &new)
	if err != nil {
		return nil, errs.New("failed to unmarshal game: %s", err)
	}

	return new, nil
}
