package server

import (
	"context"
	"log"
	"testing"

	"github.com/dylanlott/edh-go/persistence"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func Test_graphQLServer_Signup(t *testing.T) {
	type args struct {
		username string
		password string
	}
	tests := []struct {
		name string
		args args
		want *User
	}{
		{
			name: "should create a user and persist it to the appDB.",
			args: args{
				username: "shakezula",
				password: "ohhellyeah",
			},
			want: &User{
				Username: "shakezula",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, err := persistence.NewAppDatabase("../persistence/migrations/")
			if err != nil {
				t.Errorf("failed to create persistence: %s", err)
			}

			cardDB, err := persistence.NewSQLite("../persistence/AllPrintings.sqlite")
			if err != nil {
				t.Errorf("failed to create cardDB: %s", err)
			}

			s, err := NewGraphQLServer(nil, db, cardDB)
			if err != nil {
				t.Errorf("failed to start server: %s", err)
			}
			got, err := s.Signup(context.Background(), tt.args.username, tt.args.password)
			if diff := cmp.Diff(got, tt.want, cmpopts.IgnoreFields(User{}, "ID")); diff != "" {
				log.Printf("DIFF: %s", diff)
				t.Errorf("wanted: %+v - got: %+v", tt.want, got)
			}
		})
	}
}
