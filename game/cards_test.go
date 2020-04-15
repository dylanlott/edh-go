package game

import (
	"fmt"
	"testing"

	"github.com/dylanlott/edh-go/persistence"

	"github.com/stretchr/testify/assert"
)

func TestNewDeckList(t *testing.T) {
	db, err := persistence.NewSQLite("../persistence/mtgallcards.sqlite")
	assert.NoError(t, err)
	list, errors := NewDecklist(db, testdata)
	assert.Equal(t, 0, len(errors))
	assert.Equal(t, 2, len(list))

	t.Run("test shuffle on deck", func(t *testing.T) {
		deck := list
		shuffled, err := deck.Shuffle()
		assert.NoError(t, err)
		assert.Equal(t, 2, len(shuffled))
	})

	t.Run("test draw", func(t *testing.T) {
		deck := list
		fmt.Printf("deck length: %d", len(deck))
		card, shuffled := Draw(deck, 1)
		assert.Equal(t, 1, len(shuffled))
		assert.Equal(t, 1, len(deck))
		assert.NotNil(t, card)
		t.Logf("drew card: %+v\n", card)
	})

	t.Run("test fetch", func(t *testing.T) {
		deck := list
		c := Card{
			Name: "Warlord's Fury",
		}
		Fetch(c, deck)
	})
}

func TestQuery(t *testing.T) {
	db, err := persistence.NewSQLite("../persistence/mtgallcards.sqlite")
	assert.NoError(t, err)
	assert.NotNil(t, db)

	t.Run("fetch single card", func(t *testing.T) {
		t.Skip()
	})
}

const testdata = `Warlord's Fury
Teysa, Envoy of Ghosts`
