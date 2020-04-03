package persistence

import (
	"testing"
)

func TestNewRedis(t *testing.T) {
	config := make(Config)
	r, err := NewRedis(config)
	if err != nil {
		t.Logf("FAILED: %s", err)
		t.Fail()
	}

	if r == nil {
		t.Logf("FAILED: Redis can't be nil - %+v", r)
		t.Fail()
	}

	val, ok, err := r.Get("testkey")
	if err != nil {
		t.Logf("failed to get testkey")
		t.Fail()
	}

	t.Logf("val: %+v", val)

	// it shouldn't exist yet, so it should return false
	if ok {
		t.Fail()
	}

}
