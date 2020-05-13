//go:generate go run github.com/99designs/gqlgen
package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/99designs/gqlgen/handler"
	"github.com/go-redis/redis"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
	"github.com/segmentio/ksuid"
	"github.com/tinrab/retry"

	"github.com/dylanlott/edh-go/persistence"
)

type contextKey string

// Observer must be fulfilled for anything that's listening to the events that
// come off of a game. GraphQL mutations trigger these events, which get
// pushed out to the rest of the users in the game.
// These are pure functions for a reason - channels can get incredibly messy
// and event emitters can already be very complicated.
type Observer interface {
	Joined(ctx context.Context, game *Game) (*Game, error)
	Updated(ctx context.Context, game *Game) (*Game, error)
	Errored(ctx context.Context, game *Game, err error)
}

const (
	userContextKey = contextKey("user")
)

// graphQLServer binds the whole app together.
type graphQLServer struct {
	mutex sync.Mutex

	// TODO: Remove redisClient and switch over to our Persistence interface
	redisClient *redis.Client

	// Directory maps game ID's to a Game pointer
	Directory map[string]*Game

	// Persistence layers
	kv     persistence.Persistence
	db     persistence.Database
	cardDB persistence.Database

	// TODO: Remove these channel implementations and work our own observer
	// pattern in
	messageChannels map[string]chan *Message
	userChannels    map[string]chan string
	observers       []Observer
}

// NewGraphQLServer creates a new server to attach the database, game engine,
// and graphql connections together
func NewGraphQLServer(
	kv persistence.KV,
	appDB persistence.Database,
	cardDB persistence.Database,
) (*graphQLServer, error) {
	// TODO: Remove this redis client and wire chat up to KV interface instead
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	retry.ForeverSleep(2*time.Second, func(_ int) error {
		_, err := client.Ping().Result()
		if err != nil {
			log.Printf("error connecting to redis: %+v\n", err)
		}
		return err
	})
	return &graphQLServer{
		cardDB:          cardDB,
		db:              appDB,
		kv:              kv,
		redisClient:     client,
		messageChannels: map[string]chan *Message{},
		userChannels:    map[string]chan string{},
		// TODO: Add Game channels here
		// gameChannels: map[string]chan *Game{},
		mutex: sync.Mutex{},
	}, nil
}

func (s *graphQLServer) Serve(route string, port int) error {
	mux := http.NewServeMux()
	mux.Handle(
		route,
		handler.GraphQL(NewExecutableSchema(Config{Resolvers: s}),
			handler.WebsocketUpgrader(websocket.Upgrader{
				CheckOrigin: func(r *http.Request) bool {
					return true
				},
			}),
		),
	)
	mux.Handle("/playground", handler.Playground("GraphQL", route))
	log.Println("serving graphiql at localhost:8080/playground")

	handler := cors.AllowAll().Handler(mux)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), handler)
}

func (s *graphQLServer) PostMessage(ctx context.Context, user string, text string) (*Message, error) {
	err := s.createUser(user)
	if err != nil {
		return nil, err
	}

	// Create message
	m := &Message{
		ID:        ksuid.New().String(),
		CreatedAt: time.Now().UTC(),
		Text:      text,
		User:      user,
	}
	mj, _ := json.Marshal(m)

	// Update messages to key off of `message:<game_id>` and `message:<room_id>`
	if err := s.redisClient.LPush("messages", mj).Err(); err != nil {
		log.Println(err)
		return nil, err
	}
	// Notify new message
	s.mutex.Lock()
	for _, ch := range s.messageChannels {
		ch <- m
	}
	s.mutex.Unlock()
	return m, nil
}

func (s *graphQLServer) Messages(ctx context.Context) ([]*Message, error) {
	cmd := s.redisClient.LRange("messages", 0, -1)
	if cmd.Err() != nil {
		log.Println(cmd.Err())
		return nil, cmd.Err()
	}
	res, err := cmd.Result()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	messages := []*Message{}
	for _, mj := range res {
		m := &Message{}
		err = json.Unmarshal([]byte(mj), &m)
		if err != nil {
			log.Printf("error unmarhsaling json Message: %s", err)
		}
		messages = append(messages, m)
	}
	return messages, nil
}

func (s *graphQLServer) Users(ctx context.Context) ([]string, error) {
	// TODO: Make this persist via SQLite
	cmd := s.redisClient.SMembers("users")
	if cmd.Err() != nil {
		log.Println(cmd.Err())
		return nil, cmd.Err()
	}
	res, err := cmd.Result()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return res, nil
}

func (s *graphQLServer) MessagePosted(ctx context.Context, user string) (<-chan *Message, error) {
	err := s.createUser(user)
	if err != nil {
		return nil, err
	}

	// Create new channel for request
	messages := make(chan *Message, 1)
	s.mutex.Lock()
	s.messageChannels[user] = messages
	s.mutex.Unlock()

	// Delete channel when done
	go func() {
		<-ctx.Done()
		s.mutex.Lock()
		delete(s.messageChannels, user)
		s.mutex.Unlock()
	}()

	return messages, nil
}

func (s *graphQLServer) UserJoined(ctx context.Context, user string) (<-chan string, error) {
	err := s.createUser(user)
	if err != nil {
		return nil, err
	}

	// Create new channel for request
	users := make(chan string, 1)
	s.mutex.Lock()
	s.userChannels[user] = users
	s.mutex.Unlock()

	// Delete channel when done
	go func() {
		<-ctx.Done()
		s.mutex.Lock()
		delete(s.userChannels, user)
		s.mutex.Unlock()
	}()

	return users, nil
}

func (s *graphQLServer) createUser(user string) error {
	// Upsert user
	if err := s.redisClient.SAdd("users", user).Err(); err != nil {
		return err
	}
	// Notify new user joined
	s.mutex.Lock()
	for _, ch := range s.userChannels {
		ch <- user
	}
	s.mutex.Unlock()
	return nil
}

func (s *graphQLServer) Mutation() MutationResolver {
	return s
}

func (s *graphQLServer) Query() QueryResolver {
	return s
}

func (s *graphQLServer) Subscription() SubscriptionResolver {
	return s
}
