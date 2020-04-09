package game

import (
	"errors"
	"sync"
	"time"

	"github.com/zeebo/errs"

	"github.com/dylanlott/edh-go/persistence"
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

// Counter is a general type of Counter on any Card or Player.
type Counter int

// UserID is used for external routing and relation to Users when we go live.
// It has validation and authorization methods assigned to it.
type UserID string

// GameID is a string that uniquely identifies a Game through out the entire
// system. This Game tracks all of the players and the board state alterations
// of each, as well as metadata around each game.
type GameID string

// Game maintains a Game state with mutexes for protection
type Game struct {
	sync.Mutex

	DB persistence.Persistence

	Name      string
	ID        GameID
	StartTime time.Time
	Players   map[UserID]*PlayerState
}

// PlayerState maintains a state for each player that is mutex protected.
type PlayerState struct {
	sync.Mutex

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
	// Revealed	 CardList

	// How should we account for other players taking control of cards?
	// There are lots of control effects in MTG, having a visual
	// representation of this control would be beneficial.

	// Counters include all game effects on Player
	Data map[string]Counter
}

var _ = (FullGame)(&Game{})

// NewGame creates a new Game object to manipulate the game board state.
func NewGame(players map[UserID]Deck, db persistence.Persistence) (*Game, error) {
	p := make(map[UserID]*PlayerState)

	for userID, decklist := range players {
		if len(decklist.Cards) != 99 {
			return nil, errs.New("deck must have exactly 99 cards")
		}

		if len(decklist.Commander) > 1 {
			return nil, errs.New("must have only one commander")
		}

		if userID == "" {
			return nil, errs.New("userID must not be empty")
		}

		p[userID] = &PlayerState{
			PlayerID:  userID,
			DB:        db,
			Library:   decklist.Cards,
			Commander: decklist.Commander,
			Graveyard: CardList{},
			Exiled:    CardList{},
			Field:     CardList{},
		}
	}

	// TODO: Persist the game state at creation time.

	g := &Game{
		Players: p,
	}
	return g, errs.New("failed to create new game")
}

// Returns the player state for a playerID.
func (g *Game) Get(player UserID) (*PlayerState, error) {
	return g.Players[player], nil
}

// Joins a player to a a game. If no game exists, it will create one.
func (g *Game) Join(deck Deck, player UserID) error {
	return errors.New("not impl")
}

// Leave removes a player from a Game.
func (g *Game) Leave(player UserID) error {
	return errors.New("not impl")
}
