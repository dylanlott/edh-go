package main

import (
	"errors"
	"sync"
)

// FullGame implements all of the below code in a neat wrapper
type FullGame interface {
	Shuffle(player UserID) (Deck, error)
	Add(card Card, player UserID) error
	Remove(card Card, player UserID) error
	Join(deck Deck, player UserID) error
	Leave(player UserID) error
}

// Actions represents a change in the board state.
type Action struct{}

// Card tracks the properties of a Card in a given Game
type Card struct {
	Name     string
	Counters map[string]int
	Details  map[string]string
}

// UserID is used for external routing and relation to Users when we go live
type UserID string

// Game maintains a Game state with mutexes for protection
type Game struct {
	sync.Mutex

	Name    string
	ID      string
	Players map[UserID]PlayerState
}

// PlayerState maintains a state for each player that is mutex protected.
type PlayerState struct {
	sync.Mutex

	Commander CardList
	Hand      CardList
	Library   CardList
	Graveyard CardList
	Exiled    CardList

	// Counters include all game effects on Player
	Counters map[string]int
}

// Deck tracks a list of Cards for a Game. These are ephemeral and
// are created and destroyed per game.
type Deck struct {
	Name   string
	Cards  CardList
	Player string // refers to player ID, only for internet purposes
}

// NewGame creates a new Game object to manipulate the game board state.
func NewGame(players map[UserID]PlayerState) FullGame {
	return &Game{}
}

// Shuffles a Deck
func (g *Game) Shuffle(player UserID) (Deck, error) {
	// deck := g.Decks[player]
	return Deck{}, errors.New("not impl")
}

func (g *Game) Remove(card Card, player UserID) error {
	return errors.New("not impl")
}

func (g *Game) Add(card Card, player UserID) error {
	return errors.New("not impl")
}

func (g *Game) Join(deck Deck, player UserID) error {
	return errors.New("not impl")
}

func (g *Game) Leave(player UserID) error {
	return errors.New("not impl")
}
