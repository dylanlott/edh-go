package persistence

// Value is a type for handling and validating Values in the game engine
type Value string

// Key is a type for handling and validating Keys in the game engine
type Key string

// Persistence allows us to use any database for persisting the state of this
// app if it can fulfill the Persistence layer.
type Persistence interface {
	Put(key Key, val Value) (Value, error)
	Get(key Key) (Value, bool, error)
}
