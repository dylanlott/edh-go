package server

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/zeebo/errs"
	"golang.org/x/crypto/bcrypt"
)

// CreateTokenEndpoint ...
func (s *graphQLServer) Signup(ctx context.Context, username string, password string) (*User, error) {
	hashed, err := hashPassword(password)
	if err != nil {
		return nil, errs.Wrap(err)
	}
	id := uuid.New().String()
	stmt := `
	INSERT INTO users (uuid, username, password)
	VALUES ($1, $2, $3);
	`
	_, err = s.db.Exec(stmt, id, username, hashed)
	if err != nil {
		return nil, errs.Wrap(err)
	}

	return &User{
		ID:       id,
		Username: username,
	}, nil
}

func (s *graphQLServer) Login(ctx context.Context, username string, password string) (*User, error) {
	return nil, errors.New("not impl")
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
