package game

import (
	"testing"

	"github.com/dylanlott/edh-go/persistence"
	"github.com/stretchr/testify/assert"
)

func TestDecks(t *testing.T) {
	if !testing.Short() {
		// NB: these will fail without database access
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
			drawn, library, err := Draw(deck, 1)
			assert.NoError(t, err)
			assert.NotNil(t, drawn)
			assert.NotNil(t, library)
			assert.Equal(t, 3, len(library))
			assert.Equal(t, 1, len(drawn))

			// 3 cards in deck, draw all 3
			// This should not fail because drawing your entire library is OK,
			drawn, library, err = Draw(library, 3)
			assert.NoError(t, err)
			assert.NotNil(t, library)
			assert.NotNil(t, drawn)
			assert.Equal(t, 3, len(drawn))

			// 0 cards left in deck, draw 1, this should fail.
			drawn, library, err = Draw(library, 1)
			assert.Error(t, err)
			assert.Nil(t, drawn)
			assert.Nil(t, library)
			assert.EqualError(t, err, "check yourself before you deck yourself")
		})

		t.Run("test fetch", func(t *testing.T) {
			deck := makeDeck(t)
			c := Card{
				Name: "Warlord's Fury",
			}
			card, list, err := Fetch(c, deck)
			assert.NoError(t, err)
			assert.NotNil(t, card)
			assert.NotNil(t, list)
		})

		t.Run("test put", func(t *testing.T) {
			deck := makeDeck(t)
			cards := CardList{
				{
					Name: "Goblin Warlord",
				},
			}
			lib, err := Put(deck, cards, 0, false)
			assert.NoError(t, err)
			assert.Equal(t, 5, len(lib))

			drawn, lib, err := Draw(lib, 1)
			assert.NoError(t, err)
			assert.NotNil(t, drawn)
			assert.NotNil(t, lib)
			assert.Equal(t, drawn[0].Name, "Goblin Warlord")
		})
	}
}

func makeDeck(t *testing.T) CardList {
	if testing.Short() {
		return CardList{
			Card{
				Name: "Teysa, Envoy of Ghosts",
			},
			Card{
				Name: "Warlord's Fury",
			},

			Card{
				Name: "Karlov of the Ghost Council",
			},

			Card{
				Name: "Shock",
			},
		}
	} else {
		db, err := persistence.NewSQLite("../persistence/mtgallcards.sqlite")
		if err != nil {
			t.Skipf("unable to connect to db - skipping database tests")
		}
		list, errors := NewDecklist(db, testdata)
		assert.Equal(t, 0, len(errors))
		assert.Equal(t, 4, len(list))

		return list
	}
}

// TODO: Make test cases handle dual cards with `//` in the name
// E.g. Expansion / Explosion, Insult / Injury, etc...

const testdata = `Warlord's Fury
Teysa, Envoy of Ghosts
Shock
Karlov of the Ghost Council
`
