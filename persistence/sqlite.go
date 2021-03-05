package persistence

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
	"github.com/zeebo/errs"
)

var (
	// Directory used for loading migrations
	migrationsDir = "file://./persistence/migrations/"
)

// NewAppDatabase returns a migrated app database or an error
func NewAppDatabase(path string, migrationsDir string) (*DB, error) {
	wrapped, err := NewSQLite(path)
	if err != nil {
		return nil, errs.Wrap(err)
	}

	migrated, err := applyMigrations(wrapped.db, migrationsDir)
	if err != nil {
		return nil, errs.Wrap(err)
	}
	log.Printf("successfully migrated sqlite3 db: %+v", migrated)
	return &DB{
		db: migrated,
	}, nil
}

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
	// return db
	return &DB{
		db: db,
	}, nil
}

func applyMigrations(db *sql.DB, migrationsDir string) (*sql.DB, error) {
	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		log.Printf("failed to create db instance: %s", err)
		return nil, errs.Wrap(err)
	}

	m, err := migrate.NewWithDatabaseInstance(migrationsDir, "sqlite3", driver)
	if err != nil {
		log.Printf("failed to create migration with database instance: %s", err)
		return nil, errs.Wrap(err)
	}
	err = m.Up()
	if err != nil {
		if err != migrate.ErrNoChange {
			log.Printf("failed to migrate db: %s", err)
			return nil, errs.Wrap(err)
		}
	}

	return db, nil
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
