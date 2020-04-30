package server

import (
	"context"

	"github.com/zeebo/errs"
)

func (s *graphQLServer) Signup(ctx context.Context, input *InputSignup) (*User, error) {
	return nil, errs.New("not impl")
}
