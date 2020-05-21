package server

import (
	"context"

	"github.com/zeebo/errs"
)
func (s *graphQLServer) Cards(ctx context.Context, id *string, name *string) ([]*Card, error) {
  return nil, errs.New("not impl")
}
