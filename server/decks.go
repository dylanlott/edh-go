package server

import (
	"context"
	"fmt"

	"github.com/zeebo/errs"
)

func (s *graphQLServer) Decks(ctx context.Context, userID string) ([]*Deck, error) {
	fmt.Printf("getting decks for user: %s", userID)
	rows, err := s.db.Query(`SELECT * FROM decks WHERE userID = ?`, userID)
	if err != nil {
		return nil, errs.Wrap(err)
	}

	for rows.Next() {

	}

	return nil, errs.New("not impl")
}
