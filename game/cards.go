package game

import (
	sdk "github.com/MagicTheGathering/mtg-sdk-go"
	"github.com/zeebo/errs"
)

// Card tracks the properties of a Card in a given Game
type Card struct {
	Name string

	// Track counters on a card
	Counters map[string]Counter

	// wrappers around the mtg sdk card api
	CardInfo sdk.Card
	ID       sdk.CardId
}

// Query will try to find card info for a given CardID with MTG SDK
func (c Card) Query() {}

// CardList exposes a set of methods for manipulating a list of Cards
type CardList []Card

// Deck is the top level resource for a given Deck
type Deck struct {
	Name      string
	Commander CardList
	Cards     CardList
	Owner     UserID
}

// Creates a new decklist from a line delimited list of card names.
func NewDecklist(decklist string) (CardList, error) {
	return CardList{}, errs.New("not impl")
}

// Shuffle is a sugar method to make Shuffling a list of Cards easier.
func (c CardList) Shuffle() (CardList, error) {
	return []Card{}, errs.New("not impl")
}

// Validate will valiate the CardList against the format specified in args.
func (c CardList) Validate(format string) bool {
	switch format {
	case "commander":
		break
	case "modern":
		break
	case "standard":
		break
	default:
		return false
	}

	return false
}

// Fetch removes a card from the library and puts into the player's Hand
func (c CardList) Fetch(card Card) (CardList, error) {
	// check if card is in deck
	// remove it if it is, and put it into player's Hand instead.
	// return the new card list or an error
	return nil, errs.New("not impl")
}

// Returns the top card of the Deck into the player's Hand
func (c CardList) Draw() Card {
	return Card{}
}

// TODO: Implement the go Sort interface on Cards here for sorting methods
func (c CardList) Sort() error {
	return errs.New("not impl")
}
