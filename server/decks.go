package server

import (
	"context"

	"github.com/zeebo/errs"
)

func (s *graphQLServer) Decks(ctx context.Context, userID string) ([]*Deck, error) {
	return nil, errs.New("not impl")
}
