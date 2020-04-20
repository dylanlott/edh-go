package main

import (
	"log"

	"github.com/dylanlott/edh-go/persistence"
)

func main() {
	db, err := persistence.NewSQLite("./persistence/mtgallcards.sqlite")
	if err != nil {
		log.Fatalf("failed to open SQLite card database: %s", err)
	}
	log.Printf("started SQLite database: %+v\n", db)
}
