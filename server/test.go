package server

import (
	"os"
	"testing"

	"github.com/dylanlott/edh-go/persistence"
)

func testAPI(t *testing.T) *graphQLServer {
	// Update these as you need to but don't commit any changes to this file.
	cfg := Conf{
		RedisURL:    "redis://localhost:6379",
		PostgresURL: "postgres://edhgo:edhgodev@localhost:5432/edhgo?sslmode=disable",
		DefaultPort: 8080,
	}
	path, err := os.Getwd()
	if err != nil {
		t.Errorf("failed to find homedir: %s", err)
	}
	t.Logf("test server path: %+v", path)
	cardDB, err := persistence.NewSQLite("../persistence/AllPrintings.sqlite")
	if err != nil {
		t.Errorf("failed to open cardDB for games_test: %s", err)
	}
	kv, err := persistence.NewRedis("redis://localhost:6379", "", persistence.Config{})
	if err != nil {
		t.Errorf("failed to get kv from redis: %s", err)
	}

	appDB, err := persistence.NewAppDatabase("../persistence/migrations/", cfg.PostgresURL)
	if err != nil {
		t.Errorf("failed to get migrated app instance: %s", err)
	}

	s, err := NewGraphQLServer(kv, appDB, cardDB, cfg)
	if err != nil {
		t.Errorf("failed to create new test server: %+v", err)
	}
	return s
}
