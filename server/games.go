package server

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/zeebo/errs"
)

func (s *graphQLServer) Games(ctx context.Context) ([]*Game, error) {
	games := []*Game{}
	for _, game := range s.Directory {
		games = append(games, game)
	}
	return games, nil
}

func (s *graphQLServer) Boardstate(ctx context.Context, gameID string) ([]*BoardState, error) {
	game, ok := s.Directory[gameID]
	if !ok {
		return nil, errs.New("game does not exist with ID of %s", gameID)
	}

	return game.Players, nil
}

func (s *graphQLServer) BoardUpdate(ctx context.Context, bs InputBoardState) (<-chan *Game, error) {
	fmt.Printf("boardupdate: %+v\n", bs)
	game, ok := s.Directory[bs.GameID]
	if !ok {
		return nil, errs.New("game does not exist with ID of %s", bs.GameID)
	}

	s.mutex.Lock()
	// take new boardstate and persist it, and update it
	player := bs.UserID
	for i, p := range game.Players {
		if p.User.ID == player {
			// swap old player `p` state with new player state `bs`
			// TODO: Add user boardstate channel and push boardstate across channels
			game.Players[i] = convertInputBoardState(bs)
		}
	}
	s.mutex.Unlock()

	return s.gameChannels[game.ID], errs.New("not impl")
}

// UpdateGame is what's used to change the name of the game, format, insert
// or remove players, or change other meta informatin about a game.
// Game shouldn't touch BoardState information ever.
func (s *graphQLServer) UpdateGame(ctx context.Context, inputGame InputGame) (*Game, error) {
	return nil, errs.New("not impl")
}

// createGame is untested currently
func (s *graphQLServer) CreateGame(ctx context.Context, inputGame InputGame) (*Game, error) {
	g := &Game{
		ID:        uuid.New().String(),
		CreatedAt: time.Now(),
		Players:   []*BoardState{},
	}

	// TODO: Flesh out boardstate
	// TODO: Make boardstate persist
	for _, player := range inputGame.Players {
		bs := &BoardState{
			User: &User{
				ID:       uuid.New().String(), // used for storing user items in redis
				Username: player.Username,
			},
			GameID: g.ID,
		}
		g.Players = append(g.Players, bs)
	}

	// Set game in directory for access
	s.mutex.Lock()
	s.gameChannels[g.ID] = make(chan *Game, 1)
	s.Directory[g.ID] = g
	s.mutex.Unlock()

	// Alert observers
	for _, obs := range s.observers {
		_, err := obs.Joined(ctx, g)
		if err != nil {
			log.Printf("error alerting observers: %+v", err)
		}
	}

	return g, nil
}

func (s *graphQLServer) UpdateBoardState(ctx context.Context, boardstate InputBoardState) (*BoardState, error) {
	pushBoardStateUpdate(ctx, s.observers, boardstate)
	return nil, errs.New("not impl")
}
func pushBoardStateUpdate(ctx context.Context, observers []Observer, input InputBoardState) {
	for _, obs := range observers {
		fmt.Printf("observer: %+v\n", obs)
		fmt.Printf("board state updated: %+v\n", input)
	}
}

func convertInputBoardState(bs InputBoardState) *BoardState {
	return nil
}
