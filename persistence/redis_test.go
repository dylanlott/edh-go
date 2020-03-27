package persistence

import (
	"testing"
)

func TestNewRedis(t *testing.T) {
	r, err := NewRedis()
	if err != nil {
		t.Logf("FAILED: %s", err)
		t.Fail()
	}

	if r == nil {
		t.Logf("FAILED: Redis can't be nil - %+v", r)
		t.Fail()
	}

	key := []byte("key1")
	v, ok, err := r.Get(key)
	if err != nil {
		t.Fail()
	}
	if v == nil && ok {
		t.Fail()
	}

	p, err := r.Put([]byte("key2"), []byte("1"))
	if err != nil {
		t.Fail()
	}

	if p == nil {
		t.Fail()
	}
}
