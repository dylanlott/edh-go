package game

import (
	"sync"

	"github.com/zeebo/errs"

	"github.com/dylanlott/edh-go/persistence"
)

// PlayerState maintains a state for each player that is mutex protected.
type PlayerState struct {
	sync.Mutex

	// gameID is the GameID of the game the player is currently in.
	GameID GameID

	// playerID assigns a unique playerID to this board state
	PlayerID UserID

	// get a reference to the database for persistencea
	DB persistence.Persistence

	BoardState BoardState
}

// BoardState holds the board state for a given PlayerState.
type BoardState struct {
	Commander CardList
	Partner   CardList
	Hand      CardList
	Library   CardList
	Graveyard CardList
	Exiled    CardList
	Field     CardList
	Counters  map[string]Counter
	Revealed  CardList
	Stolen    CardList
}

// NewPlayer returns a new PlayerState
func NewPlayer(gameID GameID, player UserID, deck Deck, db persistence.Persistence) *PlayerState {
	return &PlayerState{
		GameID:     gameID,
		PlayerID:   player,
		BoardState: BoardState{},
	}
}

/*
 * Persitence Interactions
 */

// Persist will save the player's board state to Persistence.
// This operation is thread safe and acquires a lock on the player struct.
func (p *PlayerState) Persist() error {
	// This is inefficient because it goes through each property and saves it
	return errs.New("not impl")
}

// PlayerState() returns the *PlayerState or an error. This operation
// requires a read lock and is memory safe.
func (p *PlayerState) PlayerState() (*PlayerState, error) {
	return nil, errs.New("not impl")
}

// Deck Methods
func (p *PlayerState) AddToBattlefield(card Card) (*PlayerState, error) {
	return nil, errs.New("not impl")
}

func (p *PlayerState) RemoveFromBattlefield(cards CardList)   {}
func (p *PlayerState) AddCounters(updates map[string]Counter) {}
func (p *PlayerState) Reveal(cards CardList)                  {}
func (p *PlayerState) Draw(library CardList)                  {}
func (p *PlayerState) Shuffle(cards CardList)                 {}
func (p *PlayerState) Discard(card CardList)                  {}
func (p *PlayerState) DiscardAtRandom(card CardList)          {}
func (p *PlayerState) Fetch(card Card)                        {}
func (p *PlayerState) AddToLibrary(card Card, pos int)        {}
func (p *PlayerState) AddToGraveyard(card Card)               {}
func (p *PlayerState) AddToExile(card Card)                   {}
func (p *PlayerState) Scry(num int)                           {}

// Move will move a single Card{} around from one list to another using Fetch.
// TODO: This function will need to be fleshed out
func Move(from CardList, to CardList, card Card) (listFrom, listTo CardList, err error) {
	list, fetched, err := Fetch(card, from)
	if err != nil {
		return nil, nil, errs.Wrap(err)
	}

	destination := CardList{}
	destination = append(destination, to...)
	destination = append(destination, fetched)

	return list, destination, nil
}
