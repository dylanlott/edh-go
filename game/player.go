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
	gameID GameID
	// playerID assigns a unique playerID to this board state
	PlayerID UserID

	// get a reference to the database for persistencea
	DB persistence.Persistence

	Commander CardList
	Partner   CardList
	Hand      CardList
	Library   CardList
	Graveyard CardList
	Exiled    CardList
	Field     CardList

	// This is for generally revealing cards to opponents.
	Revealed CardList

	// How should we account for other players taking control of cards?
	// There are lots of control effects in MTG, having a visual
	// representation of this control would be beneficial.

	// Counters include all game effects on Player
	Data map[string]Counter
}

// Persitence Interactions

// Persist will save the player's board state to Persistence.
// This operation is thread safe and acquires a lock on the player struct.
func (p *PlayerState) Persist() error {
	return errs.New("not impl")
}

// PlayerState() returns the *PlayerState or an error. This operation
// requires a read lock and is memory safe.
func (p *PlayerState) PlayerState() (*PlayerState, error) {
	return nil, errs.New("not impl")
}

// Deck Methods
func (p *PlayerState) CreateDeck(cards CardList)              {}
func (p *PlayerState) AddToBattlefield(cards CardList)        {}
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
