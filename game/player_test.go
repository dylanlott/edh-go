package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayerBoardStates(t *testing.T) {
	player := NewPlayer("game_id", "player_id", Deck{}, nil)
	assert.NotNil(t, player)
	assert.Equal(t, player.GameID, GameID("game_id"))
	assert.Equal(t, player.PlayerID, UserID("player_id"))

	t.Run("test Move", func(t *testing.T) {
	})

	t.Run("test add card to battlefield", func(t *testing.T) {
	})

	t.Run("test remove from battlefield", func(t *testing.T) {
	})

	t.Run("test add counters", func(t *testing.T) {
	})

	t.Run("test reveal", func(t *testing.T) {
	})

	t.Run("test draw", func(t *testing.T) {
	})

	t.Run("test shuffle", func(t *testing.T) {
	})
}
