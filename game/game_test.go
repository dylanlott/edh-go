package game

import (
	"testing"

	"github.com/dylanlott/edh-go/persistence"
)

// TestNewFullGame tries to start a redis instance and uses it to run an
// integration test suite.
func TestNewFullGame(t *testing.T) {
	// players default starting state
	players := make(map[UserID]Deck)

	players["player1"] = Deck{
		Name:  "Karlov Voltron",
		Cards: CardList{},
	}

	db, err := persistence.NewRedis(persistence.Config{})
	if err != nil {
		t.Logf("failed to start redis - %s - skipping tests", err)
		t.Skip()
	}
	if db == nil {
		t.Log("db was nil")
		t.Fail()
	}
}

func TestBoardState(t *testing.T) {
	players := make(map[UserID]Deck)

	// TODO: Create a deck for this test.
	players["play1"] = Deck{
		Name:  "Karlov Voltron",
		Cards: CardList{},
	}

	t.Logf("players map: %+v\n", players)

	db, err := persistence.NewRedis(persistence.Config{})
	if err != nil {
		t.Logf("failed to connect to redis - skipping tests")
		t.Skip()
	}
	t.Logf("connected to testredis: [ %+v ]", db)
}
