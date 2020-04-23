package server

import (
	"context"
	"log"

	"github.com/zeebo/errs"

	"github.com/dylanlott/edh-go/game"
)

func (s *graphQLServer) Games(ctx context.Context) ([]*Game, error) {
	return nil, errs.New("not impl")
}

func (s *graphQLServer) Boardstate(ctx context.Context) ([]*BoardState, error) {
	return nil, errs.New("not impl")
}

// createGame is untested currently
func (s *graphQLServer) createGame(user string, rawdeck string) (*Game, error) {
	// game.NewDecklist()
	players := make(map[game.UserID]game.Deck)
	playerID := game.UserID(user)
	players[playerID] = game.Deck{}
	g, err := game.NewGame(players, s.kv)
	if err != nil {
		log.Printf("error creating new game: %+v\n", err)
		return nil, errs.Wrap(err)
	}

	log.Printf("created game: %+v\n", g)

	// if err := s.redisClient.SAdd("games", game).Err(); err != nil {
	// 	log.Printf("error adding game to redis: %s", err)
	// 	return nil, err
	// }

	// for _, obs := range s.observers {
	// TODO: Context needs to be pulled through from top level of the app
	// obs.Joined(context.TODO(), g)
	// }

	return nil, errs.New("not impl")
}

// updates a Game boardstate for a given User
func (s *graphQLServer) updateGame(user string, game *Game) (*Game, error) {
	// save updated game state to redis
	// if successful, emit game updated event
	// return Game or error
	return nil, errs.New("not impl")
}

func (s *graphQLServer) UpdateBoard(ctx context.Context, user string) (*BoardState, error) {
	return nil, errs.New("not impl")
}
