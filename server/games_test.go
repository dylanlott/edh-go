package server

import (
	"context"
	"testing"

	"github.com/dylanlott/edh-go/persistence"
)

func TestCreateGame(t *testing.T) {
	var cases = []struct {
		name  string
		input *InputCreateGame
		want  interface{}
		err   error
	}{
		{
			name: "happy path creation",
			input: &InputCreateGame{
				Players: []*InputBoardState{},
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			s := getNewServer(t)
			result, err := s.CreateGame(context.Background(), *tt.input)
			if err != nil {
				if tt.err != err {
					t.Errorf("undesired error: %+v", err)
				}
			}
			if result.ID == "" {
				t.Errorf("games must have an ID")
			}
			if len(result.PlayerIDs) != len(tt.input.Players) {
				t.Errorf("failed to add correct amount of players to the game")
			}
		})
	}
}

func getNewServer(t *testing.T) *graphQLServer {
	cardDB, err := persistence.NewSQLite("../persistence/db.sqlite")
	if err != nil {
		t.Errorf("failed to open cardDB for games_test: %s", err)
	}

	appDB, err := persistence.NewSQLite("../persistence/db.sqlite")
	if err != nil {
		t.Errorf("failed to open appDB for games_test: %s", err)
	}

	s, err := NewGraphQLServer(nil, appDB, cardDB)
	if err != nil {
		t.Errorf("failed to create new test server: %+v", err)
	}

	return s
}
