package game

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dylanlott/edh-go/persistence"
)

// TestNewFullGame tries to start a redis instance and uses it to run an
// integration test suite.
func TestNewFullGame(t *testing.T) {
	players := make(map[UserID]Deck)
	players["player1"] = Deck{
		Name: "Karlov Voltron",
		Commander: CardList{
			{Name: "Karlov of the Ghost Council"},
		},
		Cards: TestDeck,
	}

	db, err := persistence.NewRedis(persistence.Config{})
	assert.NoError(t, err)
	assert.NotNil(t, db)
}

func TestBoardState(t *testing.T) {
	db, err := persistence.NewRedis(persistence.Config{})
	assert.NoError(t, err)

	players := make(map[UserID]Deck)
	players["player1"] = Deck{
		Name: "Karlov Voltron",
		Commander: CardList{
			{Name: "Karlov of the Ghost Council"},
		},
		Cards: TestDeck,
	}

	g, err := NewGame(players, db)
	assert.NoError(t, err)
	assert.NotNil(t, g)

	t.Logf("game created: %+v", g)
}

// TestDeck is used for testing purposes. it's a dummy deck of nothing but swamps.
var TestDeck = CardList{
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
	{Name: "Swamp"},
}
