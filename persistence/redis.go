package persistence

import (
	"github.com/zeebo/errs"

	"github.com/go-redis/redis/v7"
)

// redisDB implelement Persistence with the Redis driver
type redisDB struct {
	client *redis.Client
}

type Config map[string]string

// NewRedis returns a new Redis Persistence that can be used
// in the application to persist and update state.
func NewRedis(config Config) (*redisDB, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return &redisDB{
		client: client,
	}, nil
}

// Put willj insert a value into the DB
func (r *redisDB) Put(key Key, val Value) (Value, error) {
	k, err := key.String()
	if err != nil {
		return Value(""), errs.New("failed to get string value for key: %s", err)
	}

	err = r.client.Set(k, []byte(val), 0).Err()
	if err != nil {
		return Value(""), errs.Wrap(err)
	}

	return val, nil
}

// Get will return a value from the database from a given Key.
// Keys need to be formatted correctly
func (r *redisDB) Get(key Key) (Value, bool, error) {
	k, err := key.String()
	if err != nil {
		return Value(""), false, errs.Wrap(err)
	}

	val, err := r.client.Get(k).Result()
	if err != nil {
		return Value(""), false, errs.New("error getting key from redis client: %s", err)
	}

	return Value(val), true, nil
}

// Do runs a redigo-style Do command through the Key Value store. This is
// generally for use with Redis commands.
func (r *redisDB) Do(cmd string, args ...interface{}) (interface{}, error) {
	return nil, errs.New("not impl")
}

// String returns the string of Value.
func (v Value) String() (string, error) {
	return string(v), nil
}

// String returns the string of the Key
func (k Key) String() (string, error) {
	return string(k), nil
}
