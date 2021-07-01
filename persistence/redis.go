package persistence

import (
	"github.com/go-redis/redis/v7"
	"github.com/zeebo/errs"
)

// redisDB implelement Persistence with the Redis driver
type redisDB struct {
	client *redis.Client
}

type config struct {
	RedisURL string `envconfig:"REDIS_URL"`
}

type Config map[string]string

// NewRedis returns a new Redis Persistence that can be used
// in the application to persist and update state.
func NewRedis(url string, pass string, opts Config) (*redisDB, error) {
	if url == "" {
		// use default
		url = "redis://localhost:6379"
	}

	client := redis.NewClient(&redis.Options{
		Addr:     url,
		Password: pass, // no password set
		DB:       0,    // use default DB
	})

	return &redisDB{
		client: client,
	}, nil
}

// Put will insert a value into the DB
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

// String returns the string of Value.
func (v Value) String() (string, error) {
	return string(v), nil
}

// String returns the string of the Key
func (k Key) String() (string, error) {
	return string(k), nil
}
