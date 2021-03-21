package server

import (
	"context"
	"fmt"

	"github.com/zeebo/errs"
)

// CreateDeck is used when a Game has been created and a User is adding a Deck to the BoardState
func (s *graphQLServer) CreateDeck(ctx context.Context, deck *InputDeck) (*BoardState, error) {
	if deck == nil {
		return nil, errs.New("must input a deck ")
	}

	// Create immutable BoardState object
	bs := &BoardState{}

	for _, card := range deck.Cards {
		fmt.Printf("adding card to deck: %+v\n", card)
	}
	// NB: We need to verify the length of the returned deck.
	// If there isn't 99 cards, we need to figure out why or throw an error.
	// NB: Handle split cards with Split property in DB?

	// parse card names from input deck (if not already in separated strings)
	// fetch cards from database
	// add to the library
	// add commander to the list
	// add format into boardstate

	return bs, errs.New("not impl")
}
