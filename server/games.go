package server

import (
	"context"
	"fmt"

	"github.com/zeebo/errs"
)

func (s *graphQLServer) Games(ctx context.Context) ([]*Game, error) {
	return nil, errs.New("not impl")
}

func (s *graphQLServer) Boardstate(ctx context.Context) ([]*BoardState, error) {
	return nil, errs.New("not impl")
}

// createGame is untested currently
func (s *graphQLServer) CreateGame(ctx context.Context, inputGame *InputGame) (*Game, error) {
	for _, player := range inputGame.Players {
		fmt.Printf("player: %+v\n", player)
	}

	// if err := s.redisClient.SAdd("games", game).Err(); err != nil {
	// 	log.Printf("error adding game to redis: %s", err)
	// 	return nil, err
	// }

	for _, obs := range s.observers {
		// TODO: Context needs to be pulled through from top level of the app
		// TODO: Make sure to pass the correct, fleshed out Game here.
		obs.Joined(context.TODO(), &Game{})
	}

	return &Game{}, errs.New("not impl")
}

func (s *graphQLServer) UpdateBoardState(ctx context.Context, boardstate InputBoardState) (*BoardState, error) {
	return nil, errs.New("not impl")
}

func (s *graphQLServer) CreateDeck(ctx context.Context, inputDeck *InputDeck) (*Deck, error) {
	return nil, errs.New("not impl")
}