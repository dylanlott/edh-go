package server

import (
	"context"
	"testing"

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
				Token:    nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := testAPI(t)
			got, err := s.Signup(context.Background(), tt.args.username, tt.args.password)
			if diff := cmp.Diff(got, tt.want, cmpopts.IgnoreFields(User{}, "ID")); diff != "" {
				t.Errorf("graphQLServer.Signup: wanted: %+v - got: %+v", tt.want, got)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("graphQLServer.Signup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want, cmpopts.IgnoreFields(User{}, "ID")); diff != "" {
				t.Errorf("failed to get correct user back")
			}
			if got.ID == "" {
				t.Errorf("failed to set UUID on user")
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
				ctx:      context.Background(),
				username: "shakezula",
				password: "testpassword",
			},
			want: &User{
				Username: "shakezula",
			},
			wantErr: false,
		},
		{
			name: "should return an error when user not found",
			args: args{
				ctx:      context.Background(),
				username: "notfound",
				password: "testpassword",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := testAPI(t)
			created, err := s.Signup(tt.args.ctx, "shakezula", "testpassword")
			if err != nil {
				t.Errorf("failed to create user for login tests: %s", err)
			}
			got, err := s.Login(tt.args.ctx, tt.args.username, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("graphQLServer.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil {
				if got == nil {
					t.Errorf("failed to return user")
					return
				}
				if got.ID == "" {
					t.Errorf("failed to set ID")
				}
				if got.Token == nil {
					t.Errorf("failed to set token")
				}
				if got.Username != created.Username {
					t.Errorf("failed to set correct username")
				}
			}
		})
	}
}
