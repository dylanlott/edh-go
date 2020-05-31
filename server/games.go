package server

import (
	"context"
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

func (s *graphQLServer) Boardstates(ctx context.Context, gameID string, userID *string) ([]*BoardState, error) {
	game, ok := s.Directory[gameID]

	// if no userID provided, send all board states
	if userID == nil {
		if !ok {
			return nil, errs.New("game does not exist with ID of %s", gameID)
		}

		return game.Players, nil
	}

	// userID not nil, so send only that boardstate
	for _, player := range game.Players {
		if player.User.Username == *userID {
			return []*BoardState{player}, nil
		}
	}

	return nil, errs.New("no user with ID of %s found", *userID)
}

func (s *graphQLServer) GameUpdated(ctx context.Context, game InputGame) (<-chan *Game, error) {
	_, ok := s.gameChannels[game.ID]
	if !ok {
		return nil, errs.New("game does not exist with ID of %s", game.ID)
	}

	games := make(chan *Game, 1)
	s.mutex.Lock()
	log.Printf("emitting updated game event: %+v\n", game)
	s.gameChannels[game.ID] = games
	s.mutex.Unlock()

	go func() {
		<-ctx.Done()
		s.mutex.Lock()
		delete(s.gameChannels, game.ID)
		s.mutex.Unlock()
	}()

	return games, nil
}

func (s *graphQLServer) BoardUpdate(ctx context.Context, bs InputBoardState) (<-chan *BoardState, error) {
	_, ok := s.boardStates[bs.User.Username]
	if !ok {
		return nil, errs.New("no boardstate exists for that user: %s", bs.User.Username)
	}

	boardstates := make(chan *BoardState, 1)
	s.mutex.Lock()
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
// NB: Game _can_ touch boardstate right now, and it probably shouldn't.
func (s *graphQLServer) UpdateGame(ctx context.Context, inputGame InputGame) (*Game, error) {
	updated := gameFromInput(inputGame)
	s.mutex.Lock()
	s.gameChannels[inputGame.ID] <- updated
	s.mutex.Unlock()
	return updated, nil
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
				Username: player.User.Username,
			},
			GameID: g.ID,
		}
		g.Players = append(g.Players, bs)
		// assign boardstates to directory
		s.mutex.Lock()
		log.Printf("pushing boardstate to boardStates[%s]: %+v\n", player.User.Username, bs)
		// instantiate player boardstate channel for updates
		// NB: Username's must be unique.
		s.boardStates[player.User.Username] = make(chan *BoardState, 1)
		log.Printf("pushed player boardstate successfully: %+v\n", player)
		s.mutex.Unlock()
	}

	// Set game in directory for access
	s.mutex.Lock()
	log.Printf("setting gameID in directory")
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

	return g, nil
}

func (s *graphQLServer) UpdateBoardState(ctx context.Context, bs InputBoardState) (*BoardState, error) {
	updated := boardStateFromInput(bs)
	s.mutex.Lock()
	s.boardStates[bs.User.Username] <- updated
	s.mutex.Unlock()
	pushBoardStateUpdate(ctx, s.observers, bs)
	return updated, nil
}

func pushBoardStateUpdate(ctx context.Context, observers []Observer, input InputBoardState) {
	for _, obs := range observers {
		log.Printf("observers being notified: %+v\n", obs)
		log.Printf("board state updated: %+v\n", input)
	}
}

func gameFromInput(game InputGame) *Game {
	players := []*BoardState{}
	for _, p := range game.Players {
		players = append(players, boardStateFromInput(*p))
	}
	out := &Game{
		ID:     game.ID,
		Handle: game.Handle,
		Turn: &Turn{
			Player: game.Turn.Player,
			Phase:  game.Turn.Phase,
			Number: game.Turn.Number,
		},
		Players: players,
	}

	return out
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
