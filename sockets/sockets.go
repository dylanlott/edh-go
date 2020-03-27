package sockets

import (
	"fmt"

	"github.com/dylanlott/edh-go/game"

	"github.com/googollee/go-socket.io"
	"github.com/zeebo/errs"
)

type Socketer interface {
	JoinGame(gameID string, playerID game.UserID, deck game.CardList) error
	LeaveGame(gameID string, playerID string) error
}

type socketWrapper struct {
	s *socketio.Server
}

// NewSocketLayer returns an implementation of the Socketer
// interface for use in the game database.
func NewSocketLayer() (Socketer, error) {
	server, err := socketio.NewServer(nil)
	if err != nil {
		return nil, errs.Wrap(err)
	}

	// connected
	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID())
		return nil
	})

	// disconnected
	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
	})

	return &socketWrapper{
		s: server,
	}, nil
}

// JoinGame will join a player to a game and return an error if there are any
// issues.
func (s *socketWrapper) JoinGame(gameID string, playerID game.UserID, deck game.CardList) error {
	return errs.New("not impl")
}

// LeaveGame will remove a player from a game
func (s *socketWrapper) LeaveGame(gameID string, playerID string) error {
	return errs.New("not impl")
}
