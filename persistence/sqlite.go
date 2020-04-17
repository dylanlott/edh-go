package persistence

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/zeebo/errs"
)

func NewSQLite(path string) (*DB, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		fmt.Printf("error opening database: %+v\n", err)
		return nil, errs.Wrap(err)
	}

	return &DB{
		db: db,
	}, nil
}

// Query will return a *sql.Rows or an error from the database.
// NB: This needs to handle it's clean up with rows.Close()
func (db *DB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return db.db.Query(query, args...)
}

// Exec will run a statement and return the result or an error.
// NB: Exec does not need cleanup.
func (db *DB) Exec(query string, args ...interface{}) (sql.Result, error) {
	return db.db.Exec(query, args...)
}

// Ping will return an error if it can't ping the database.
func (db *DB) Ping() error {
	return db.db.Ping()
}

// Stats returns a struct of DBStats, which is mostly used for debugging
// and performance monitoring.
func (db *DB) Stats() sql.DBStats {
	return db.db.Stats()
}
