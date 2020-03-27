package game

import (
	"fmt"

	"github.com/zeebo/errs"
)

// GameKey holds a reference to the global state of a game
type GameKey string

// CardKey holds a key that ties to values for a given card that
// resides in a Game. This is one level of nesting lower on a GameKey.
type CardKey string

// Field is the generic name for different fields that need to be
// tracked for a given Player and Game
type Field string

func NewGameKey(gameID GameID, userID UserID, fieldID Field) GameKey {
	return fmt.Sprintf("%s:%s:%s", gameID, userID, fieldID)
}

func NewCardKey(gameId, playerId, cardId, fieldId string) CardKey {
	return fmt.Sprintf("%s:%s:%s:%s", gameId, playerId, cardId, fieldId)
}

func GameKeyFromString(key string) (Game, error) {
	return nil, errs.New("not impl")
}

func CardFieldFromString(key string) (CardField, error) {
	return nil, errs.New("not impl")
}
