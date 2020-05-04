package persistence

import (
	"database/sql"
)

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

// KV is the KV store for the game engine to work with.
type KV interface {
	Put(key Key, val Value) (Value, error)
	Get(key Key) (Value, bool, error)
	Do(cmd string, args ...interface{}) (interface{}, error)
}

// Database must be fulfilled for the cards package to operate correctly.
// This is mostly used for mocking out tests and simply fulfills the basic
// `database/sql` interface. Query and Exec are the main methods.
type Database interface {
	Prepare(query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
	Ping() error
	Stats() sql.DBStats
}

type JSONStorage interface {
	Put(key Key, val struct{}) (struct{}, error)
	Get(key Key) (struct{}, error)
}

// Force DB to fulfill Database
var _ = (Database)(&DB{})

// DB holds a reference to the database for internal use.
type DB struct {
	db *sql.DB
}
