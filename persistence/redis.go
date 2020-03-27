package persistence

import (
	"github.com/zeebo/errs"
)

type Value []byte
type Key []byte

type Persistence interface {
	Put(key Key, val Value) (Value, error)
	Get(key Key) (Value, bool, error)
}

// sqlite implements Persistence with the SQLite driver
type sqlite struct {
}

// redis implelement Persistence with the Redis driver
type redis struct {
}

// NewRedis returns a new Redis Persistence that can be used
// in the application to persist and update state.
func NewRedis() (Persistence, error) {
	return &redis{}, errs.New("not implemented")
}

func (r *redis) Put(key Key, val Value) (Value, error) {
	return nil, errs.New("not implemented")
}

func (r *redis) Get(key Key) (Value, bool, error) {
	return nil, false, errs.New("not implemented")
}

func (v Value) String() (string, error) {
	return "", errs.New("not implemented")
}
