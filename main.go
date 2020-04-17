package main

import (
	"log"

	"github.com/dylanlott/edh-go/server"

	"github.com/dylanlott/edh-go/persistence"
)

func main() {
	s, err := server.NewGraphQLServer("localhost:6379")
	if err != nil {
		log.Fatal(err)
	}

	db, err := persistence.NewSQLite("./persistence/mtgallcards.sqlite")
	if err != nil {
		log.Fatalf("failed to open SQLite card database: %s", err)
	}
	log.Printf("started SQLite database: %+v\n", db)

	err = s.Serve("/graphql", 8080)
	if err != nil {
		log.Fatal(err)
	}
}
