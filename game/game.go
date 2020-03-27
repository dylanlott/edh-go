package game

import (
	"errors"
	"sync"
	"time"
)

// FullGame implements all of the below code in a neat wrapper
type FullGame interface {
	// Get returns a pointer to a PlayerState
	Get(player UserID) (*PlayerState, error)

	// Join will add a player to the Game
	Join(deck Deck, player UserID) error

	// Leave will remove a player from a Game
	Leave(player UserID) error
}

// UserID is used for external routing and relation to Users when we go live
type UserID string

// Game maintains a Game state with mutexes for protection
type Game struct {
	sync.Mutex

	Name      string
	ID        string
	StartTime time.Time
	Players   map[UserID]PlayerState
}

// PlayerState maintains a state for each player that is mutex protected.
type PlayerState struct {
	sync.Mutex

	Commander CardList
	Hand      CardList
	Library   CardList
	Graveyard CardList
	Exiled    CardList
	Field     CardList

	// Counters include all game effects on Player
	Counters map[string]int
}

// NewGame creates a new Game object to manipulate the game board state.
func NewGame(players map[UserID]Deck) FullGame {
	g := &Game{}
	return g
}

// Shuffles a Deck
func (g *Game) Shuffle(player UserID) CardList {
	return nil
}

func (g *Game) Get(player UserID) (*PlayerState, error) {
	return nil, errors.New("not impl")
}

func (g *Game) Join(deck Deck, player UserID) error {
	return errors.New("not impl")
}

func (g *Game) Leave(player UserID) error {
	return errors.New("not impl")
}
