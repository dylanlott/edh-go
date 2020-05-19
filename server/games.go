package server

import (
	"context"
	"fmt"
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

func (s *graphQLServer) BoardUpdate(ctx context.Context, bs InputBoardState) (<-chan []*BoardState, error) {
	game, ok := s.Directory[bs.GameID]
	if !ok {
		return nil, errs.New("game does not exist with ID of %s", bs.GameID)
	}

	fmt.Printf("found game: %+v\n", game)

	return nil, errs.New("not impl")
}

// createGame is untested currently
func (s *graphQLServer) CreateGame(ctx context.Context, inputGame *InputGame) (*Game, error) {
	fmt.Printf("create game hit: %+v\n", inputGame)

	g := &Game{
		ID:        uuid.New().String(),
		CreatedAt: time.Now(),
		Players:   []*BoardState{},
	}

	for _, player := range inputGame.Players {
		fmt.Printf("player: %+v\n", player)
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
	s.Directory[g.ID] = g

	// Alert observers
	for _, obs := range s.observers {
		obs.Joined(ctx, g)
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
