package server

import (
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

	// prepare server after everything is connected
	// - clear redis before each test for clear environment
	// - TODO: clear postgresql before returning TestAPI
	_, err = s.redisClient.Do("FLUSHALL").Result()
	if err != nil {
		t.Fatalf("failed to flush redis for tests: %s", err)
	}

	return s
}