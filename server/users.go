package server

import (
	"context"
	"errors"
	"net/http"
)

// ValidateJWT ...
func ValidateJWT(t string) (interface{}, error) {
	return nil, errors.New("not impl")
}

// CreateTokenEndpoint ...
func CreateTokenEndpoint(response http.ResponseWriter, request *http.Request) {}

var jwtSecret []byte = []byte("thepolyglotdeveloper")

func (s *graphQLServer) Signup(ctx context.Context, username string, password string) (*User, error) {
	return nil, errors.New("not impl")
}

func (s *graphQLServer) Login(ctx context.Context, username string, password string) (*User, error) {
	return nil, errors.New("not impl")
}
