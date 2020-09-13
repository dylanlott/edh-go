package main

import (
	"fmt"
	"log"

	"github.com/kelseyhightower/envconfig"
	"github.com/zeebo/errs"

	"github.com/dylanlott/edh-go/persistence"
	"github.com/dylanlott/edh-go/server"
)

type config struct {
	RedisURL string `envconfig:"REDIS_URL"`
}

func main() {
	var cfg config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	db, err := persistence.NewSQLite("./persistence/db.sqlite")
	if err != nil {
		log.Fatal(err)
	}

	cardDB, err := persistence.NewSQLite("./persistence/AllPrintings.sqlite")
	if err != nil {
		log.Fatalf(errs.Wrap(err).Error())
	}

	s, err := server.NewGraphQLServer(nil, db, cardDB)
	if err != nil {
		log.Fatal(err)
	}

	err = s.Serve("/graphql", 8080)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("server listening on localhost:%d", 8080)
	fmt.Printf("serving graphiql playground at localhost:8080/playground")
}
