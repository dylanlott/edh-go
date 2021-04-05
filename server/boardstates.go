package server

import (
	"context"
	"fmt"
	"log"

	"github.com/zeebo/errs"
)

// BoardStateKey formats a board state key for boardstate to user mapping.
func BoardStateKey(gameID, username string) string {
	return fmt.Sprintf("%s:%s", gameID, username)
}

// Boardstates queries Redis for different boardstates per player or game
func (s *graphQLServer) Boardstates(ctx context.Context, gameID string, username *string) ([]*BoardState, error) {
	game, ok := s.Directory[gameID]
	if game == nil {
		log.Printf("Game is nil: %s %+v", gameID, s.Directory)
		return nil, errs.New("game does not exist")
	}
	if !ok {
		log.Printf("game is !ok: %+v", s.Directory)
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

// BoardUpdate returns a channel that emits all the Boardstate's over it and then
// listens for ctx.Done and then cleans up after itself.
func (s *graphQLServer) BoardUpdate(ctx context.Context, bs InputBoardState) (<-chan *Game, error) {
	// Make a boardstates channel to emit all the events on, and assign it to
	// the user who submitted to the update.
	boardstates := make(chan *BoardState, 1)
	s.mutex.Lock()
	s.boardChannels[bs.User.Username] = boardstates
	s.mutex.Unlock()

	go func() {
		<-ctx.Done()
		s.mutex.Lock()
		delete(s.boardChannels, bs.User.Username)
		s.mutex.Unlock()
	}()

	game, ok := s.Directory[bs.GameID]
	if !ok {
		return nil, errs.New("game %s does not exist", bs.GameID)
	}

	games := make(chan *Game, 1)
	games <- game
	return games, nil
}

func (s *graphQLServer) UpdateBoardState(ctx context.Context, bs InputBoardState) (*BoardState, error) {
	updated, err := boardStateFromInput(bs)
	if err != nil {
		log.Printf("UpdateBoardState failed to marshal input board state correctly: %s", err)
		return nil, fmt.Errorf("failed to marshal input boardstate: %s", err)
	}
	boardKey := BoardStateKey(bs.GameID, bs.User.Username)
	err = s.Set(boardKey, updated)
	if err != nil {
		log.Printf("error updating boardstate in redis: %s", err)
	}

	s.mutex.Lock()
	s.boardChannels[bs.User.Username] <- updated
	s.mutex.Unlock()
	return updated, nil
}
