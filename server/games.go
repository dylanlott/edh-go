package server

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/zeebo/errs"
)

func (s *graphQLServer) Games(ctx context.Context) ([]*Game, error) {
	games := []*Game{}
	for _, game := range s.Directory {
		games = append(games, game)
	}
	return games, nil
}

func (s *graphQLServer) Boardstate(ctx context.Context, gameID string) ([]*BoardState, error) {
	game, ok := s.Directory[gameID]
	if !ok {
		return nil, errs.New("game does not exist with ID of %s", gameID)
	}

	return game.Players, nil
}

func (s *graphQLServer) GameUpdated(ctx context.Context, gameID string) (<-chan *Game, error) {
	return nil, errs.New("Not impl")
}

func (s *graphQLServer) BoardUpdate(ctx context.Context, bs InputBoardState) (<-chan *BoardState, error) {
	_, ok := s.boardStates[bs.User.Username]
	if !ok {
		return nil, errs.New("no boardstate exists for that user: %s", bs.User.Username)
	}

	boardstates := make(chan *BoardState, 1)
	s.mutex.Lock()
	fmt.Printf("hitting BoardUpdate channel")
	s.boardStates[bs.User.Username] = boardstates
	s.mutex.Unlock()

	go func() {
		<-ctx.Done()
		s.mutex.Lock()
		delete(s.boardStates, bs.User.Username)
		s.mutex.Unlock()
	}()

	return boardstates, nil
}

// UpdateGame is what's used to change the name of the game, format, insert
// or remove players, or change other meta informatin about a game.
// Game shouldn't touch BoardState information ever.
func (s *graphQLServer) UpdateGame(ctx context.Context, inputGame InputGame) (*Game, error) {
	return nil, errs.New("not impl")
}

// createGame is untested currently
func (s *graphQLServer) CreateGame(ctx context.Context, inputGame InputGame) (*Game, error) {
	g := &Game{
		ID:        uuid.New().String(),
		CreatedAt: time.Now(),
		Players:   []*BoardState{},
	}

	for _, player := range inputGame.Players {
		bs := &BoardState{
			User: &User{
				ID:       uuid.New().String(),
				Username: player.Username,
			},
			GameID: g.ID,
		}
		g.Players = append(g.Players, bs)
		// assign boardstates to directory
		s.mutex.Lock()
		fmt.Printf("pushing boardstate to boardStates[%s]: %+v\n", player.Username, bs)
		// instantiate player boardstate channel for updates
		s.boardStates[player.Username] = make(chan *BoardState, 1)
		fmt.Printf("pushed boardstate successfully")
		s.mutex.Unlock()
	}

	// Set game in directory for access
	s.mutex.Lock()
	fmt.Printf("setting gameID in directory")
	s.gameChannels[g.ID] = make(chan *Game, 1)
	s.Directory[g.ID] = g
	s.mutex.Unlock()

	// Alert observers
	for _, obs := range s.observers {
		_, err := obs.Joined(ctx, g)
		if err != nil {
			log.Printf("error alerting observers: %+v", err)
		}
	}

	fmt.Printf("returning game; %+v\n", g)
	return g, nil
}

func (s *graphQLServer) UpdateBoardState(ctx context.Context, bs InputBoardState) (*BoardState, error) {
	fmt.Printf("hitting update board state")
	updated := boardStateFromInput(bs)
	s.mutex.Lock()
	fmt.Printf("updating boardstate: %+v\n", bs)
	s.boardStates[bs.User.Username] <- updated
	s.mutex.Unlock()
	fmt.Printf("returning updated; %+v\n", updated)
	return updated, nil
}

func pushBoardStateUpdate(ctx context.Context, observers []Observer, input InputBoardState) {
	for _, obs := range observers {
		fmt.Printf("observer: %+v\n", obs)
		fmt.Printf("board state updated: %+v\n", input)
	}
}

func boardStateFromInput(bs InputBoardState) *BoardState {
	out := &BoardState{
		User: &User{
			Username: bs.User.Username,
		},
		GameID: bs.GameID,
	}

	for _, c := range bs.Commander {
		out.Commander = append(out.Commander, &Card{
			Name: c.Name,
		})
	}

	for _, c := range bs.Library {
		out.Library = append(out.Library, &Card{
			Name: c.Name,
		})
	}

	for _, c := range bs.Exiled {
		out.Exiled = append(out.Exiled, &Card{
			Name: c.Name,
		})
	}

	for _, c := range bs.Field {
		out.Field = append(out.Field, &Card{
			Name: c.Name,
		})
	}

	for _, c := range bs.Hand {
		out.Hand = append(out.Hand, &Card{
			Name: c.Name,
		})
	}

	for _, c := range bs.Controlled {
		out.Controlled = append(out.Controlled, &Card{
			Name: c.Name,
		})
	}

	for _, c := range bs.Revealed {
		out.Revealed = append(out.Revealed, &Card{
			Name: c.Name,
		})
	}

	return out
}
