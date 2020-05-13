package server

import (
	"context"

	"github.com/zeebo/errs"
)

func (s *graphQLServer) Decks(ctx context.Context, userID string) ([]*Deck, error) {
	return nil, errs.New("not impl")
}

func (s *graphQLServer) CreateDeck(ctx context.Context, inputDeck *InputDeck) (*Deck, error) {
	return nil, errs.New("not impl")
}
