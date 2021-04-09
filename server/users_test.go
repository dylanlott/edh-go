package server

import (
	"context"
	"fmt"
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
		name    string
		args    args
		want    *User
		wantErr bool
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
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := newServer(t)
			got, err := s.Signup(context.Background(), tt.args.username, tt.args.password)
			if diff := cmp.Diff(got, tt.want, cmpopts.IgnoreFields(User{}, "ID")); diff != "" {
				t.Errorf("graphQLServer.Signup: wanted: %+v - got: %+v", tt.want, got)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("graphQLServer.Signup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_graphQLServer_Login(t *testing.T) {
	type args struct {
		ctx      context.Context
		username string
		password string
	}
	tests := []struct {
		name    string
		args    args
		want    *User
		wantErr bool
	}{
		{
			name: "should log in a user and give them a token",
			args: args{
				username: "shakezula",
				password: "password",
			},
			want: &User{
				Username: "shakezula",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := newServer(t)
			log.Printf("SERVER: %+v", s)
			log.Printf("SERVER db: %+v", s.db)
			got, err := s.Login(tt.args.ctx, tt.args.username, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("graphQLServer.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got,
				cmpopts.IgnoreFields(User{}, "ID", "Token"),
				cmpopts.IgnoreUnexported(User{}),
			); diff != "" {
				t.Errorf("Login: wanted: %+v - got: %+v", tt.want, got)
			}
		})
	}
}
func newServer(t *testing.T) *graphQLServer {
	host := "localhost"
	port := 5432
	user := "edhgo"
	password := "edhgodev"
	dbname := "edhgo"
	localPG := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	cfg := Conf{
		PostgresURL: localPG,
	}
	db, err := persistence.NewAppDatabase("../persistence/migrations/", cfg.PostgresURL)
	if err != nil {
		t.Errorf("failed to create persistence: %s", err)
	}

	kv, err := persistence.NewRedis("localhost:6379", "", persistence.Config{})
	if err != nil {
		t.Errorf("failed to create a KV instance: %s", err)
	}

	cardDB, err := persistence.NewSQLite("../persistence/AllPrintings.sqlite")
	if err != nil {
		t.Errorf("failed to create cardDB: %s", err)
	}

	s, err := NewGraphQLServer(kv, db, cardDB, cfg)
	if err != nil {
		t.Errorf("failed to start server: %s", err)
	}
	return s
}
