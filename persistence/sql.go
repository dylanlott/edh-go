package persistence

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/mattn/go-sqlite3"
	"github.com/zeebo/errs"
)

var (
	// Directory used for loading migrations
	defaultMigrationsDir = "file://./persistence/migrations/"
	defaultPGURL         = "postgres://edhgo:edhgodev@localhost:5432/edhgo?sslmode=disable"
)

// NewAppDatabase returns a migrated app database or an error
func NewAppDatabase(migdir, dbURL string) (*DB, error) {
	pg, err := NewPostgres(migdir, dbURL)
	if err != nil {
		log.Printf("failed to get postgres: %s", err)
		return nil, errs.Wrap(err)
	}
	return &DB{
		db: pg,
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

	log.Printf("[DB] database connection established: %+v\n", db)
	// return db
	return &DB{
		db: db,
	}, nil
}

// NewPostgres returns a migrated sql.DB with a Postgres database connection
// migdir is the relative path to the migrations directory.
func NewPostgres(migdir string, dbURL string) (*sql.DB, error) {
	// These need to directly match the Heroku environment
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Printf("failed to get new postgres: %s", err)
		return nil, errs.Wrap(err)
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Printf("failed to get postgres instance: %s", err)
		return nil, err
	}
	formattedMigrationsDir := fmt.Sprintf("file://%s", migdir)
	m, err := migrate.NewWithDatabaseInstance(formattedMigrationsDir, "postgres", driver)
	if err != nil {
		log.Printf("failed to get instance for migration: %s", err)
		return nil, errs.Wrap(err)
	}
	err = m.Up()
	if err != nil {
		log.Printf("failed to run migrations: %s", err)
		// attempt to rollback if migration fails
		return nil, errs.Wrap(m.Down())
	}

	log.Printf("returning NewPostgres db: %+v", db)
	return db, err
}

func applySqliteMigrations(db *sql.DB, migrationsDir string) (*sql.DB, error) {
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

	log.Printf("returning applyMigrations db: %+v", db)
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
	log.Printf("DB: %+v", db)
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
