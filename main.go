package main

import (
	"fmt"
	"log"

	"github.com/zeebo/errs"

	"github.com/dylanlott/edh-go/persistence"
	"github.com/dylanlott/edh-go/server"
	"github.com/kelseyhightower/envconfig"
)

func main() {
	var cfg server.Conf
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	db, err := persistence.NewAppDatabase("./persistence/migrations/", cfg.PostgresURL)
	if err != nil {
		log.Fatal(errs.Wrap(err))
	}

	cardDB, err := persistence.NewSQLite("./persistence/AllPrintings.sqlite")
	if err != nil {
		log.Fatalf(errs.Wrap(err).Error())
	}

	kv, err := persistence.NewRedis(cfg.RedisURL, "", nil)
	if err != nil {
		log.Fatalf("failed to start redis: %s", errs.Wrap(err))
	}
	s, err := server.NewGraphQLServer(kv, db, cardDB, cfg)
	if err != nil {
		log.Fatal(err)
	}

	err = s.Serve("/graphql", cfg.DefaultPort)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("serving /graphql at :%d", cfg.DefaultPort)
	fmt.Printf("serving graphiql playground at :%d/playground", cfg.DefaultPort)
}
