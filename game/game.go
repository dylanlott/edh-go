package game

import (
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"github.com/dylanlott/edh-go/persistence"
)

// FullGame implements all of the below code in a neat wrapper
type FullGame interface {
	// Get returns a pointer to a PlayerState
	Get(player UserID) (*PlayerState, error)

	// Join will add a player to the Game
	Join(deck Deck, player UserID) (*Game, error)

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

// Game maintains a Game state with mutexes for protection.
// This is where all of the socket, graphQL, and database connections get
// combined for ease of use.
type Game struct {
	sync.Mutex

	// DB holds a reference to the persistence layer so that we can always
	// have Put and Get access to the database.
	DB persistence.Database

	Name      string
	ID        GameID
	StartTime time.Time
	Players   map[UserID]*PlayerState
}

var _ = (FullGame)(&Game{})

// NewGame creates a new Game object to manipulate the game board state.
// TODO: This should be named for Commander and other formats will create
// different game types with different validations.
func NewGame(players map[UserID]Deck, db persistence.Persistence) (*Game, error) {
	p := make(map[UserID]*PlayerState)

	for userID, decklist := range players {
		// TODO: This should eventually be validated against the format being played
		// but for now this will just check 99 + 1 commander for EDH
		if len(decklist.Cards) != 99 {
			return nil, errs.New("deck must have exactly 99 cards; had %d", len(decklist.Cards))
		}

		if len(decklist.Commander) > 1 {
			return nil, errs.New("must have only one commander")
		}

		if userID == "" {
			return nil, errs.New("userID must not be empty")
		}

		p[userID] = &PlayerState{
			PlayerID:  userID,
			Library:   decklist.Cards,
			Commander: decklist.Commander,
			Graveyard: CardList{},
			Exiled:    CardList{},
			Field:     CardList{},
		}

		// TODO: Persist player state here.
		// db.SetPlayer()
	}

	gameID := uuid.New()

	g := &Game{
		ID:        GameID(gameID.String()),
		StartTime: time.Now(),
		Players:   p,
	}

	return g, nil
}

// Returns the player state for a playerID.
func (g *Game) Get(player UserID) (*PlayerState, error) {
	return g.Players[player], nil
}

// Joins a player to a a game. If no game exists, it will create one.
func (g *Game) Join(deck Deck, player UserID) (*Game, error) {
	if len(deck.Cards) != 99 {
		return nil, errs.New("deck must contain exactly 99 cards")
	}

	if len(deck.Commander) != 1 {
		return nil, errs.New("deck must have exactly one commander.")
	}

	g.Players[player] = &PlayerState{
		PlayerID:  player,
		Commander: deck.Commander,
		Library:   deck.Cards,
		Graveyard: CardList{},
		Exiled:    CardList{},
		Hand:      CardList{},
		Field:     CardList{},
	}

	return g, nil
}

// Leave removes a player from a Game.
func (g *Game) Leave(player UserID) error {
	g.Players[player] = nil
	return nil
}
