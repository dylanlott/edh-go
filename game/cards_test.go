package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDeckList(t *testing.T) {
	list, errors := NewDecklist(testdata)
	assert.Equal(t, len(errors), 0)
	assert.Equal(t, len(list), 2)
}

const testdata = `Warlord's Fury
Teysa, Envoy of Ghosts`
