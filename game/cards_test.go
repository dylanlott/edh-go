package game

import (
	"testing"

	"github.com/dylanlott/edh-go/persistence"
	"github.com/stretchr/testify/assert"
)

func TestNewDeckList(t *testing.T) {
	t.Run("test shuffle on deck", func(t *testing.T) {
		deck := makeDeck(t)
		shuffled, err := Shuffle(deck)
		assert.NoError(t, err)
		assert.Equal(t, 4, len(shuffled))
	})

	t.Run("test draw", func(t *testing.T) {
		deck := makeDeck(t)
		assert.Equal(t, 4, len(deck))

		// 4 cards in deck, draw 1
		drawn, shuffled, err := Draw(deck, 1)
		assert.NoError(t, err)
		assert.NotNil(t, drawn)
		assert.NotNil(t, shuffled)
		assert.Equal(t, 3, len(shuffled))
		assert.Equal(t, 1, len(drawn))

		// 3 cards in deck, draw all 3
		// This should not fail because drawing your entire library is OK,
		drawn, shuffled, err = Draw(shuffled, 3)
		assert.NoError(t, err)
		assert.NotNil(t, shuffled)
		assert.NotNil(t, drawn)
		assert.Equal(t, 3, len(drawn))

		// 0 cards left in deck, draw 1, this should fail.
		drawn, shuffled, err = Draw(shuffled, 1)
		assert.Error(t, err)
		assert.Nil(t, drawn)
		assert.Nil(t, shuffled)
		assert.EqualError(t, err, "check yourself before you deck yourself")
	})

	t.Run("test fetch", func(t *testing.T) {
		deck := makeDeck(t)
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

func makeDeck(t *testing.T) CardList {
	db, err := persistence.NewSQLite("../persistence/mtgallcards.sqlite")
	if err != nil {
		t.Skipf("unable to connect to db - skipping database tests")
	}
	assert.NoError(t, err)
	list, errors := NewDecklist(db, testdata)
	assert.Equal(t, 0, len(errors))
	assert.Equal(t, 4, len(list))

	return list
}

// TODO: Make test cases handle dual cards with `//` in the name
// E.g. Expansion / Explosion, Insult / Injury, etc...

const testdata = `Warlord's Fury
Teysa, Envoy of Ghosts
Shock
Karlov of the Ghost Council
`
