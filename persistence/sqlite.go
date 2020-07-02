package persistence

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/zeebo/errs"
)

// NewSQLite returns a DB object to persist data for the application.
func NewSQLite(path string) (*DB, error) {
	log.Printf("opening database connection at %s\n", path)
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Printf("error opening sqlite3 database: %+v\n", err)
		return nil, errs.Wrap(err)
	}

	err = db.Ping()
	if err != nil {
		log.Printf("error pinging database: %s", err)
		return nil, errs.Wrap(err)
	}

	log.Printf("database connection established: %+v\n", db)

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

// Prepare will return a formatted sql Stmt
func (db *DB) Prepare(query string) (*sql.Stmt, error) {
	return db.db.Prepare(query)
}
