package main

import "testing"

func TestNewFullGame(t *testing.T) {

	// players default starting state
	players := map[UserID]PlayerState{}

	players["play1"] = PlayerState{
		Commander: CardList{Card{
			Name: "Karlov of the Ghost Council",
		}},
	}
	// create a new game
	g := NewGame(players)
	d, err := g.Shuffle(UserID("play1"))
	if err != nil {
		t.Fail()
	}

	t.Logf("deck: %+v", d)
}
