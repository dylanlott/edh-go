package server

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/zeebo/errs"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret []byte = []byte("TODO:SET THIS TO FROM AN ENV VAR")

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
		log.Printf("Signup failed: %s", err)
		return nil, errs.Wrap(err)
	}

	return &User{
		ID:       id,
		Username: username,
	}, nil
}

func (s *graphQLServer) Login(ctx context.Context, username string, password string) (*User, error) {
	if password == "" {
		return nil, errs.New("must provide a password for authentication")
	}
	if username == "" {
		return nil, errs.New("must provide a username for authentication")
	}

	// check passwor
	log.Printf("attempting: %s - %s", username, password)
	log.Printf("attempting with db: %+v", s.db)
	rows, err := s.db.Query(`SELECT uuid, username, password FROM users WHERE username = $1`, username)
	if err != nil {
		log.Printf("failed to get row: %s", err)
		if errs.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		log.Printf("error Is: %s", err)
	}
	defer rows.Close()
	var user *User
	cols, _ := rows.Columns()
	log.Printf("cols: %s", cols)
	if rows == nil {
		log.Printf("ROWS IS NIL")
	}
	if err := rows.Scan(user.ID, user.Username, user.Password); err != nil {
		log.Printf("failed to scan rows into user: %s", err)
		return nil, errs.Wrap(err)
	}
	log.Printf("#USER: %+v", user)
	valid := checkPasswordHash(password, *user.Password)
	if !valid {
		return nil, errs.New("failed to authenticate")
	}
	log.Printf("valid: %+v", valid)

	// we're valid, so generate a new token and assign it to the user
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"password": user.Password,
	})
	t, error := token.SignedString(jwtSecret)
	if error != nil {
		fmt.Println(error)
	}

	// set token in redis for session comparison
	log.Printf("token generated: %s", t)
	user.Token = &t
	return user, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
