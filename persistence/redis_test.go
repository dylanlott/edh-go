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
	if ok {
		t.Fail()
	}
	if err == nil {
		t.Fail()
	}
	if val != Value("") {
		t.Logf("val: %+v", val)
		t.Fail()
	}
}
