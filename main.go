package main

import (
	"log"

	"github.com/kelseyhightower/envconfig"

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

	s, err := server.NewGraphQLServer(cfg.RedisURL)
	if err != nil {
		log.Fatal(err)
	}

	err = s.Serve("/graphql", 8080)
	if err != nil {
		log.Fatal(err)
	}
}
