package server

import (
	"context"
	"fmt"

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

func (s *graphQLServer) Boardstate(ctx context.Context, userID string) ([]*BoardState, error) {
	return nil, errs.New("not impl")
}

func (s *graphQLServer) BoardUpdate(ctx context.Context, user InputUser, bs InputBoardState) (<-chan *User, error) {
	return nil, errs.New("not impl")
}

// createGame is untested currently
func (s *graphQLServer) CreateGame(ctx context.Context, inputGame *InputGame) (*Game, error) {
	fmt.Printf("create game hit: %+v\n", inputGame)

	g := &Game{
		ID: uuid.New().String(),
	}

	for _, player := range inputGame.Players {
		fmt.Printf("player: %+v\n", player)
	}

	for _, obs := range s.observers {
		// TODO: Context needs to be pulled through from top level of the app
		// TODO: Make sure to pass the correct, fleshed out Game here.
		obs.Joined(ctx, &Game{})
	}

	// TODO: Wire Game type up to game.Game
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
