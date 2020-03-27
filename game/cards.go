package game

import "errors"

// Card tracks the properties of a Card in a given Game
type Card struct{}

// CardList exposes a set of methods for manipulating a list of Cards
type CardList []Card

// Deck is the top level resource for a given Deck
type Deck struct {
	Name      string
	Commander Card
	Cards     CardList
	Owner     UserID
}

// Shuffle is a sugar method to make Shuffling a list of Cards easier.
func (c CardList) Shuffle() (CardList, error) {
	return []Card{}, errors.New("not impl")
}

// Fetch removes a card from the library and puts into the player's Hand
func (c CardList) Fetch(card Card) (CardList, error) {
	return nil, errors.New("not impl")
}

// Returns the top card of the Deck into the player's Hand
func (c CardList) Draw() Card {
	return Card{}
}

// TODO: Implement the go Sort interface on Cards here for sorting methods
func (c CardList) Sort() error {
	return errors.New("not impl")
}
