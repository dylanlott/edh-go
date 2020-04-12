package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDeckList(t *testing.T) {
	list, errors := NewDecklist(testdata)
	assert.Equal(t, len(errors), 0)
	assert.Equal(t, len(list), 3)

	for _, card := range list {
		assert.NotEqual(t, card.Name, "")
	}
}

const testdata = `
Karlov of the Ghost Council
Teysa, Envoy of Ghosts
Alela, Artful Provocateur
`
