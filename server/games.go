package server

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/imdario/mergo"
	"github.com/zeebo/errs"
)

// IPersistence defines the persistence interface for the server.
// This interface stores Game and BoardStates for realtime interaction.
type IPersistence interface {
	Set(key string, value interface{}) error
	Get(key string, dest interface{}) error
}

var _ IPersistence = (&graphQLServer{})

// Games returns a list of Games.
func (s *graphQLServer) Games(ctx context.Context, gameID *string) ([]*Game, error) {
	if gameID == nil {
		games := []*Game{}
		for _, game := range s.Directory {
			games = append(games, game)
		}

		return games, nil
	}

	game, ok := s.Directory[*gameID]
	if !ok {
		return nil, errs.New("game [%+v] does not exist", gameID)
	}

	return []*Game{game}, nil
}

// Boardstates queries Redis for different boardstates per player or game
func (s *graphQLServer) Boardstates(ctx context.Context, gameID string, username *string) ([]*BoardState, error) {
	game, ok := s.Directory[gameID]
	if game == nil {
		return nil, errs.New("game does not exist")
	}
	if !ok {
		return nil, errs.New("game does not exist")
	}

	// if username is not provided, send all
	if username == nil {
		boardstates := []*BoardState{}
		for _, p := range game.PlayerIDs {
			board := &BoardState{}
			boardKey := BoardStateKey(game.ID, p.Username)
			err := s.Get(boardKey, &board)
			if err != nil {
				log.Printf("error fetching user boardstate from redis: %s", err)
			}
			boardstates = append(boardstates, board)
		}
		return boardstates, nil
	}

	boardstates := []*BoardState{}
	for _, p := range game.PlayerIDs {
		if p.Username == *username {
			board := &BoardState{}
			boardKey := BoardStateKey(game.ID, p.Username)
			err := s.Get(boardKey, &board)
			if err != nil {
				log.Printf("error fetching user boardstate from redis: %s", err)
			}

			boardstates = append(boardstates, board)
		}
	}

	if len(boardstates) == 0 {
		return []*BoardState{}, errs.New("no boardstate for user %s found", *username)
	}

	return boardstates, nil
}

func getUsers(players []*InputUser) []*User {
	users := []*User{}
	for _, p := range players {
		u := &User{
			ID:       *p.ID,
			Username: p.Username,
		}
		users = append(users, u)
	}

	return users
}

func getTurn(turn *InputTurn) *Turn {
	return &Turn{
		Number: turn.Number,
		Phase:  turn.Phase,
		Player: turn.Player,
	}
}

func (s *graphQLServer) GameUpdated(ctx context.Context, game InputGame) (<-chan *Game, error) {
	found, ok := s.Directory[game.ID]
	if !ok {
		return nil, errs.New("game does not exist with ID of %s", game.ID)
	}

	output := &Game{
		ID:        found.ID,
		Handle:    game.Handle,
		CreatedAt: found.CreatedAt,
		PlayerIDs: getUsers(game.PlayerIDs),
		Turn:      getTurn(game.Turn),
		Rules:     found.Rules,
	}

	games := make(chan *Game, 1)
	s.mutex.Lock()
	s.gameChannels[game.ID] = games
	// TODO: turn output into update of new and old game
	s.Directory[game.ID] = output
	s.mutex.Unlock()

	go func() {
		<-ctx.Done()
		s.mutex.Lock()
		delete(s.gameChannels, game.ID)
		s.mutex.Unlock()
	}()

	return games, nil
}

// BoardUpdate returns a channel that emits all the Boardstate's over it and then
// listens for ctx.Done and then cleans up after itself.
func (s *graphQLServer) BoardUpdate(ctx context.Context, bs InputBoardState) (<-chan *Game, error) {
	// Make a boardstates channel to emit all the events on, and assign it to
	// the user who submitted to the update.
	boardstates := make(chan *BoardState, 1)
	s.mutex.Lock()
	s.boardChannels[bs.User.Username] = boardstates
	s.mutex.Unlock()

	go func() {
		<-ctx.Done()
		s.mutex.Lock()
		delete(s.boardChannels, bs.User.Username)
		s.mutex.Unlock()
	}()

	game, ok := s.Directory[bs.GameID]
	if !ok {
		return nil, errs.New("game %s does not exist", bs.GameID)
	}

	games := make(chan *Game, 1)
	games <- game
	return games, nil
}

// UpdateGame is what's used to change the name of the game, format, insert
// or remove players, or change other meta informatin about a game.
// NB: Game _can_ touch boardstate right now, and it probably shouldn't.
// NB: We eventually need a stronger support for combining two structs of different types
// for GraphQL. Something like https://play.golang.org/p/UBCq0waIEe should eventually be used.
func (s *graphQLServer) UpdateGame(ctx context.Context, new InputGame) (*Game, error) {
	// check existence of game, fail if not found
	old, ok := s.Directory[new.ID]
	if !ok {
		return nil, errs.New("Game with ID %s does not exist", new.ID)
	}

	// update old game with new game data
	if err := mergo.Merge(&new, old); err != nil {
		return nil, errs.New("Failed to merge old game with new game: %s", err)
	}

	// cast new game into Game for GraphQL
	game := &Game{}
	if err := mergo.Merge(game, new); err != nil {
		return nil, errs.New("Failed to merge new game: %s", err)
	}

	s.mutex.Lock()
	s.Directory[new.ID] = game
	s.gameChannels[new.ID] <- game
	s.mutex.Unlock()

	return game, nil
}

// JoinGame ...
func (s *graphQLServer) JoinGame(ctx context.Context, input *InputJoinGame) (*Game, error) {
	// TODO: We check for game existence a lot, we should probably make this a function
	s.mutex.RLock()
	game, ok := s.Directory[input.ID]
	if !ok {
		return nil, errs.New("Game with ID %s does not exist", input.ID)
	}
	s.mutex.RUnlock()

	user := &User{
		ID:       *input.User.ID,
		Username: input.User.Username,
	}

	// Init default boardstate minus library and commander
	bs := &BoardState{
		User:       user,
		Life:       input.BoardState.Life,
		GameID:     game.ID,
		Hand:       getCards(input.BoardState.Hand),
		Exiled:     getCards(input.BoardState.Exiled),
		Revealed:   getCards(input.BoardState.Revealed),
		Field:      getCards(input.BoardState.Field),
		Controlled: getCards(input.BoardState.Controlled),
	}

	library, err := s.createLibraryFromDecklist(ctx, *input.Decklist)
	if err != nil {
		// Fail gracefully and still populate basic cards
		bs.Library = getCards(input.BoardState.Library)
	} else {
		// Happy path
		bs.Library = library
	}

	// TODO: This will eventually have to check the rules of the game to see if it's a
	// Commander game, but for now this works for EDH MVP.
	if len(input.BoardState.Commander) == 0 {
		return nil, errs.New("must supply a Commander for your deck.")
	}

	// TODO: Make this handle multiple commanders?
	commander, err := s.Card(ctx, input.BoardState.Commander[0].Name, nil)
	if err != nil {
		// fail gracefully and use their card name so they can still play a game
		cmdr := getCards(input.BoardState.Commander)
		bs.Commander = []*Card{cmdr[0]}
	} else {
		bs.Commander = []*Card{commander[0]}
	}

	shuff, err := Shuffle(bs.Library)
	if err != nil {
		return nil, err
	}
	bs.Library = shuff

	game.PlayerIDs = append(game.PlayerIDs, user)

	// set board state in Redis
	boardKey := BoardStateKey(game.ID, user.Username)
	err = s.Set(boardKey, bs)
	if err != nil {
		log.Printf("error persisting boardstate into redis: %s", err)
		return nil, err
	}

	s.mutex.Lock()
	s.Directory[game.ID] = game
	s.mutex.Unlock()
	return game, nil
}

// createGame is untested currently
func (s *graphQLServer) CreateGame(ctx context.Context, inputGame InputCreateGame) (*Game, error) {
	g := &Game{
		ID:        uuid.New().String(),
		CreatedAt: time.Now(),
		PlayerIDs: []*User{},
		// NB: We're only supporting EDH at this time. We will add more flexible validation later.
		Rules: []*Rule{
			{
				Name:  "format",
				Value: "EDH",
			},
			{
				Name:  "deck_size",
				Value: "99",
			},
		},
	}

	for _, player := range inputGame.Players {
		// TODO: Deck validation should happen here.
		user := &User{
			ID:       uuid.New().String(),
			Username: player.User.Username,
		}
		g.PlayerIDs = append(g.PlayerIDs, user)

		// Init default boardstate minus library and commander
		bs := &BoardState{
			User:       user,
			Life:       player.Life,
			GameID:     g.ID,
			Hand:       getCards(player.Hand),
			Exiled:     getCards(player.Exiled),
			Revealed:   getCards(player.Revealed),
			Field:      getCards(player.Field),
			Controlled: getCards(player.Controlled),
		}

		var decklist string
		if inputGame.Players[0].Decklist != nil {
			decklist = string(*inputGame.Players[0].Decklist)
		}
		library, err := s.createLibraryFromDecklist(ctx, decklist)
		if err != nil {
			// Fail gracefully and still populate basic cards
			bs.Library = getCards(player.Library)
		} else {
			// Happy path
			bs.Library = library
		}

		commander, err := s.Card(ctx, player.Commander[0].Name, nil)
		if err != nil {
			// fail gracefully and use their card name so they can still play a game
			inputCard := getCards(player.Commander)
			bs.Commander = []*Card{inputCard[0]}
		} else {
			bs.Commander = []*Card{commander[0]}
		}

		shuff, err := Shuffle(bs.Library)
		if err != nil {
			return nil, err
		}
		bs.Library = shuff
		boardKey := BoardStateKey(g.ID, bs.User.Username)
		err = s.Set(boardKey, bs)
		if err != nil {
			log.Printf("error persisting boardstate into redis: %s", err)
			return nil, err
		}

		// save the baordChannels to the same key format of <gameID:username>
		s.mutex.Lock()
		s.boardChannels[boardKey] = make(chan *BoardState, 1)
		s.mutex.Unlock()
	}

	// Set game in directory for access
	s.mutex.Lock()
	s.gameChannels[g.ID] = make(chan *Game, 1)
	s.Directory[g.ID] = g
	s.mutex.Unlock()

	// persist it to Redis
	err := s.Set(g.ID, g)
	if err != nil {
		log.Printf("error setting Game to redis: %+v\n", err)
	}

	return g, nil
}

func (s *graphQLServer) UpdateBoardState(ctx context.Context, bs InputBoardState) (*BoardState, error) {
	updated, err := BoardStateFromInput(bs)
	if err != nil {
		return nil, fmt.Errorf("failed to get boardstate from input: %s", err)
	}
	boardKey := BoardStateKey(bs.GameID, bs.User.Username)
	err = s.Set(boardKey, updated)
	if err != nil {
		log.Printf("error updating boardstate in redis: %s", err)
	}

	s.mutex.Lock()
	s.boardChannels[bs.User.Username] <- updated
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

func GameFromInput(game InputGame) (*Game, error) {
	data, err := json.Marshal(game)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal input game: %s", err)
	}

	var g *Game
	err = json.Unmarshal(data, &g)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal game from input: %s", err)
	}

	return g, nil
}

func getPlayerIDs(inputUsers []*InputUser) []*User {
	var u []*User
	for _, i := range inputUsers {
		u = append(u, &User{
			ID:       *i.ID,
			Username: i.Username,
		})
	}

	return u
}

func BoardStateFromInput(bd InputBoardState) (*BoardState, error) {
	data, err := json.Marshal(bd)
	if err != nil {
		return nil, fmt.Errorf("failed to get board state from input: %s", err)
	}

	var b *BoardState
	err = json.Unmarshal(data, &b)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal boardstate from input: %s", err)
	}

	return b, nil
}

func getCards(inputCards []*InputCard) []*Card {
	cardList := []*Card{}

	for _, card := range inputCards {
		c := &Card{
			Name: card.Name,
		}

		if card.ID != nil {
			c.ID = *card.ID
		}

		cardList = append(cardList, c)
	}

	return cardList
}

func (s *graphQLServer) createLibraryFromDecklist(ctx context.Context, decklist string) ([]*Card, error) {
	if decklist == "" {
		return []*Card{}, errs.New("must provide cards in decklist to create a library")
	}
	trimmed := strings.TrimSpace(decklist)
	r := csv.NewReader(strings.NewReader(trimmed))
	cards := []*Card{}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			// handle error path
			log.Printf("error reading record: %+v", err)
			return nil, errs.New("failed to parse CSV: %s", err)
		}

		trimmed := strings.TrimSpace(record[0])
		quantity, err := strconv.ParseInt(trimmed, 0, 64)
		if err != nil {
			// handle error
			log.Printf("error parsing quantity: %+v\n", err)
			// assume quantity = 1
			quantity = 1
		}

		// NB: In the future, this should be optimized to be one query for all the cards
		// instead of a query for each card in the deck.
		name := record[1]
		card, err := s.Card(ctx, name, nil)
		if err != nil {
			// handle lookup error
			log.Printf("error looking up card: %+v\n", err)
			cards = append(cards, &Card{
				Name: name,
			})
			continue
		}

		if card == nil {
			fmt.Printf("failed to find card: %s", name)
		}

		// happy path
		var num int64 = 1
		for num <= quantity {
			// Fail gracefully if we can't find the card
			if len(card) == 0 {
				fmt.Printf("failed to find card- adding dummy card instead")
				cards = append(cards, &Card{
					Name: name,
				})
				num++
			} else {
				// add the first card that's returned from the database
				// NB: This is going to need to be handled eventually
				cards = append(cards, card[0])
				num++
			}
		}

		continue
	}

	return cards, nil
}

// BoardStateKey formats a board state key for boardstate to user mapping.
func BoardStateKey(gameID, username string) string {
	return fmt.Sprintf("%s:%s", gameID, username)
}

// Set will set a value into the Redis client and returns an error, if any
func (s *graphQLServer) Set(key string, value interface{}) error {
	exp, err := time.ParseDuration("12h")
	if err != nil {
		exp = 0
	}
	p, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return s.redisClient.Set(key, p, exp).Err()
}

// Get returns a value from Redis client to `dest` and returns an error, if any
func (s *graphQLServer) Get(key string, dest interface{}) error {
	p, err := s.redisClient.Get(key).Result()
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(p), dest)
}
