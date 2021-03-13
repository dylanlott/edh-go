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
	RedisURL    string `envconfig:"REDIS_URL"`
	PostgresURL string `envconfig:"POSTGRES_URL"`
}

func main() {
	var cfg config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	db, err := persistence.NewAppDatabase("./persistence/migrations/")
	if err != nil {
		log.Printf("error getting app db: %s", err)
		log.Fatal(err)
	}

	cardDB, err := persistence.NewSQLite("./persistence/AllPrintings.sqlite")
	if err != nil {
		log.Fatalf(errs.Wrap(err).Error())
	}

	kv, err := persistence.NewRedis("localhost:6379", "", nil)
	s, err := server.NewGraphQLServer(kv, db, cardDB)
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
