package sockets

import (
	"fmt"

	"github.com/dylanlott/edh-go/game"
	"github.com/dylanlott/edh-go/persistence"

	socketio "github.com/googollee/go-socket.io"
	"github.com/zeebo/errs"
)

// Socketer is what every socket implementation must fulfill to work with our system.
type Socketer interface {
	// JoinGame allows us to handle joining login regardless of underlying
	// socket implementation
	JoinGame(gameID string, playerID game.UserID, deck game.CardList) (*Game, error)

	// LeaveGame allows us to handle leaving games regardless of the underlying
	// socket implementation.
	LeaveGame(gameID string, playerID string) error

	// TODO
	// GetGame returns a *Game based on the gameID that you pass it.
	// GetGame(gameID string) (*Game, error)

	// RegisterHandler allows us to pass arbitrary handlers to a socket system.
	RegisterHandler(room string, event string, handler HandlerFunc) error
}

// Config allows for config values to be passed through to the socket layer
type Config map[string]string

// HandlerFunc declares a type that any handler can use to declare listeners
// and emitters with.
// TODO: See if this needs to take anything or if this will work because
// access to the socketWrapper is available.
type HandlerFunc func() error

// SocketGame wraps a game with a database and a socket layer to give the
// socket layer access to PlayerState methods.
type SocketGame struct {
	db         persistence.Persistence
	boardstate game.Game
	socket     socketWrapper
}

// Game joins together a boardstate with a socket layer for interaction over
type Game struct {
	id string

	// TODO: Boardstate should be an interface rather than a concrete struct
	boardstate *game.Game

	socket Socketer
}

// socketWrapper holds a reference to the socketio.Server and wraps it with
// any other data necessary to manage the socket server
type socketWrapper struct {
	Client *socketio.Server
}

var _ = (Socketer)(&socketWrapper{})

// NewSocketLayer returns an implementation of the Socketer
// interface for use in the game database.
func NewSocketLayer() (*socketWrapper, error) {
	server, err := socketio.NewServer(nil)
	if err != nil {
		return nil, errs.Wrap(err)
	}

	// connected
	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected /:", s.ID())
		return nil
	})

	// disconnected
	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
	})

	// errored
	server.OnError("/", func(s socketio.Conn, err error) {
		fmt.Println("error: ", err)
	})

	return &socketWrapper{
		Client: server,
	}, nil
}

// JoinGame will join a player to a game and return an error if there are any
// issues.
func (s *socketWrapper) JoinGame(gameID string, playerID game.UserID, deck game.CardList) (*Game, error) {
	return &Game{}, errs.New("not impl")
}

// LeaveGame will remove a player from a game
func (s *socketWrapper) LeaveGame(gameID string, playerID string) error {
	return errs.New("not impl")
}

// RegisterHandler allows for registration of handlers from outside of the
// initialization of the new socket layer area. There is probably a better
// way to handle this, but this is clean enough for now.
func (s *socketWrapper) RegisterHandler(room string, event string, handler HandlerFunc) error {
	s.Client.OnEvent(room, event, handler)
	return nil
}

func (s *socketWrapper) GetClient() *socketio.Server {
	fmt.Printf("get client has been hit\n")
	return s.Client
}
