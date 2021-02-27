package server

import (
	"context"
	"errors"
)

// CreateTokenEndpoint ...
func (s *graphQLServer) Signup(ctx context.Context, username string, password string) (*User, error) {
	return nil, errors.New("not impl")
}

func (s *graphQLServer) Login(ctx context.Context, username string, password string) (*User, error) {
	return nil, errors.New("not impl")
}
