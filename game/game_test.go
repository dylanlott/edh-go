package game

import "testing"

func TestNewFullGame(t *testing.T) {
	// players default starting state
	players := make(map[UserID]Deck)

	players["player1"] = Deck{
		Name:  "Karlov Voltron",
		Cards: CardList{},
	}

	// create a new game
	g, err := NewGame(players)
	if err != nil {
		t.Fail()
	}

	deck := Deck{
		Name: "test deck",
		Cards: CardList{
			Card{
				Name: "test card 1",
			},
		},
		Commander: Card{
			Name: "Karlov of the Ghost Council",
		},
	}

	// Join Game
	err = g.Join(deck, UserID("player2"))
	if err != nil {
		t.Fail()
	}

	// Leave Game
	err = g.Leave(UserID("player2"))
	if err != nil {
		t.Fail()
	}

	// TODO: Test that getting a user that has left the game
	// returns a nil and an error to make sure that's correct

	p, err := g.Get("game1", "player1")
	if err != nil {
		t.Fail()
	}
	if p == nil {
		t.Fail()
	}

	// TODO: ensure that player left game in game state
}

func TestBoardState(t *testing.T) {
	players := make(map[UserID]Deck)

	players["play1"] = Deck{
		Name:  "Karlov Voltron",
		Cards: CardList{},
	}

	g, err := NewGame(players)
	if err != nil {
		t.Fail()
	}

	p1, err := g.Get("game1", "play1")
	if err != nil {
		t.Fail()
	}
	if p1 == nil {
		t.Fail()
	}

	t.Logf("player1: %+v", p1)

	if p1.PlayerID == "" {
		t.Fail()
	}
}
