package sockets

import (
	"fmt"

	"../game"

	"github.com/googollee/go-socket.io"
	"github.com/zeebo/errs"
)

type Socketer interface {
	JoinGame(gameID string, playerID player, deck CardList) error
	LeaveGame(gameID string, playerID string)
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
func (s *socketWrapper) JoinGame(gameID string, playerID string, deck CardList) {
	return errs.New("not impl")
}
