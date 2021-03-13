package persistence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRedis(t *testing.T) {
	r, err := NewRedis("localhost:6379", "", nil)
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

	val, err = r.Put(Key("abc"), Value("value"))
	assert.NoError(t, err)
	assert.NotNil(t, val)
	assert.Equal(t, val, Value("value"))
}
