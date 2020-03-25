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

type sqlite struct {
}

// redis fulfills the Persistence interface with an adapter
type redis struct {
}

// NewDataLayer returns a new Persistence that can be used
// in the application to persist and update state.
func NewDataLayer() (Persistence, error) {
	return &redis{}, errs.New("not implemented")
}

func (r *redis) Put(key Key, val Value) (Value, error) {
	return nil, errs.New("not implemented")
}

func (r *redis) Get(key Key) (Value, bool, error) {
	return nil, false, errs.New("not implemented")
}
