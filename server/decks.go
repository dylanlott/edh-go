package server

import (
	"context"
	"fmt"

	"github.com/zeebo/errs"
)

func (s *graphQLServer) Decks(ctx context.Context, userID string) ([]*Deck, error) {
	fmt.Printf("userID: %+v\n", userID)
	return nil, errs.New("not impl")
}
