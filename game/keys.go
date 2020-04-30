package game

import (
	"fmt"
)

// GameKey holds a reference to the global state of a game
type GameKey string

// CardKey holds a key that ties to values for a given card that
// resides in a Game. This is one level of nesting lower on a GameKey.
type CardKey string

// CardListKey holds a reference to a CardList such as a library or a graveyard
type CardListKey string

// Field is the generic name for different fields that need to be
// tracked for a given Player and Game
type Field string

func NewCardListKey(gameID GameID, userID UserID, name string) CardListKey {
	key := fmt.Sprintf("%s:%s:%s", gameID, userID, name)
	return CardListKey(key)
}

func NewGameKey(gameID GameID, userID UserID, fieldID Field) GameKey {
	key := fmt.Sprintf("%s:%s:%s", gameID, userID, fieldID)
	return GameKey(key)
}

func NewCardKey(gameId, playerId, cardId, fieldId string) CardKey {
	key := fmt.Sprintf("%s:%s:%s:%s", gameId, playerId, cardId, fieldId)
	return CardKey(key)
}

// String returns the string value of a GameKey
func (g GameKey) String() string {
	return string(g)
}

func ParseGameKey(str string) GameKey {
	return GameKey("not implemented")
}
