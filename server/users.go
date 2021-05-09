package server

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/zeebo/errs"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret []byte = []byte(os.Getenv("JWT_SECRET"))

// CreateTokenEndpoint ...
func (s *graphQLServer) Signup(ctx context.Context, username string, password string) (*User, error) {
	if password == "" {
		return nil, errs.New("must provide a password")
	}
	hashed, err := hashPassword(password)
	if err != nil {
		return nil, errs.Wrap(err)
	}
	id := uuid.New().String()
	stmt := `
	INSERT INTO "users" (uuid, username, password)
	VALUES ($1, $2, $3)
	RETURNING uuid, username;
	`
	result, err := s.db.Query(stmt, id, username, hashed)
	if err != nil {
		return nil, errs.Wrap(err)
	}
	defer result.Close()
	user := &User{}
	for result.Next() {
		if err := result.Scan(&user.ID, &user.Username); err != nil {
			return nil, errs.Wrap(err)
		}
	}

	return user, nil
}

func (s *graphQLServer) Login(ctx context.Context, username string, password string) (*User, error) {
	if password == "" {
		return nil, errs.New("must provide a password for authentication")
	}
	if username == "" {
		return nil, errs.New("must provide a username for authentication")
	}

	// TODO: We need to enforce uniqueness as a constraint on username in the DB
	q := `SELECT "uuid", "username", "password" FROM "users" WHERE username=$1`
	rows, err := s.db.Query(q, username)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Zero rows found")
			return nil, errs.New("user not found")
		} else {
			return nil, errs.Wrap(err)
		}
	}

	defer rows.Close()
	user := &User{}
	var hash string
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Username, &hash); err != nil {
			log.Printf("failed to scan in user at login: %s", err)
			return nil, errs.Wrap(err)
		}
	}

	log.Printf("Logging in user: %+v", user)
	// check password validity, return if invalid
	valid := checkPasswordHash(password, hash)
	if !valid {
		return nil, errs.New("failed to authenticate")
	}

	// we're valid, so generate a new token and assign it to the user
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"password": password,
	})

	// and attempt to sign that token with our server's jwtSecret
	t, err := token.SignedString(jwtSecret)
	if err != nil {
		return nil, errs.Wrap(err)
	}

	// set password to blank so we don't return the sensitive material
	user.Password = nil

	// set token in redis for session comparison - expires every 2 weeks
	s.redisClient.Set(user.Username, t, time.Duration(time.Hour*24*14))

	user.Token = &t

	log.Printf("returning logged in user: %+v", user)
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
