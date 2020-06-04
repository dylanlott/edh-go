package server

import (
	"context"

	"github.com/zeebo/errs"
)

// Returns a list of queryable decks. NOT IMPL RIGHT NOW.
func (s *graphQLServer) Decks(ctx context.Context, userID string) ([]*Deck, error) {
	return nil, errs.New("not impl")
}

// CreateDeck is used when a Game has been created and a User is adding a Deck to the BoardState
func (s *graphQLServer) CreateDeck(ctx context.Context, inputDeck *InputDeck) (*BoardState, error) {
	if inputDeck == nil {
		return nil, errs.New("must input a deck ")
	}

	// Create immutable BoardState object
	bs := &BoardState{}

	// TODO: Create the query with the WHERE IN for each card in the stack
	rows, err := s.cardDB.Query()
	if err != nil {
		return nil, errs.New("not impl")
	}

	for rows.Next() {
		// Do shit
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
