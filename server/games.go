package server

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/dylanlott/edh-go/persistence"
	"github.com/google/uuid"
	"github.com/zeebo/errs"
)

// IPersistence defines the persistence interface for the server.
// This interface stores Game and BoardStates for realtime interaction.
type IPersistence interface {
	Set(key string, value interface{}) error
	Get(key string, dest interface{}) error
}

// Observer is an interface that will eventually get fulfilled for
// emitting game state updates.
type Observer interface{}

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
	// TODO: Use json Marshaling hack here
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

func (s *graphQLServer) GameUpdated(ctx context.Context, from InputGame) (<-chan *Game, error) {
	_, ok := s.Directory[from.ID]
	if !ok {
		return nil, errs.New("game does not exist with ID of %s", from.ID)
	}

	// HACK: This gets around us having to do a lot of complicated
	// checking and assignment and let's the json/encoder library
	// handle the type conversion instead.
	b, err := json.Marshal(from)
	if err != nil {
		return nil, errs.New("error marshaling from: %+v", err)
	}
	game := &Game{}
	err = json.Unmarshal(b, &game)
	if err != nil {
		return nil, errs.New("failed to unmarshal into *Game: %+v", err)
	}

	games := make(chan *Game, 1)
	s.mutex.Lock()
	s.Directory[game.ID] = game
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
func (s *graphQLServer) UpdateGame(ctx context.Context, new InputGame) (*Game, error) {
	// check existence of game, fail if not found
	_, ok := s.Directory[new.ID]
	if !ok {
		return nil, errs.New("Game with ID %s does not exist", new.ID)
	}

	b, err := json.Marshal(new)
	if err != nil {
		return nil, errs.New("failed to marshal input game: %s", err)
	}
	game := &Game{}
	err = json.Unmarshal(b, &game)
	if err != nil {
		return nil, errs.New("failed to unmarshal game: %s", err)
	}

	s.mutex.Lock()
	s.Directory[new.ID] = game
	log.Printf("pushing new game on channel: %+v", game)
	s.gameChannels[new.ID] <- game
	s.mutex.Unlock()

	return game, nil
}

// JoinGame ...
func (s *graphQLServer) JoinGame(ctx context.Context, input *InputJoinGame) (*Game, error) {
	// TODO: Check context for User auth and append user info that way
	// TODO: We check for game existence a lot, we should probably make this a function
	s.mutex.RLock()
	if input.User.ID == nil {
		return nil, errors.New("must provide a user ID to join a game")
	}
	if input.User.Username == "" {
		return nil, errors.New("must provide a username to join a game")
	}

	game, ok := s.Directory[input.ID]
	if !ok {
		log.Printf("game lookup for id %s failed", input.ID)
		return nil, errs.New("Game with ID %s does not exist", input.ID)
	}
	s.mutex.RUnlock()

	user := &User{
		Username: input.User.Username,
		ID:       *input.User.ID,
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
		log.Printf("error creating library from decklist: %+v", err)
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
		log.Printf("error getting commander for deck: %+v", err)
		// fail gracefully and use their card name so they can still play a game
		cmdr := getCards(input.BoardState.Commander)
		bs.Commander = []*Card{cmdr[0]}
	} else {
		bs.Commander = []*Card{commander[0]}
	}

	shuff, err := Shuffle(bs.Library)
	if err != nil {
		log.Printf("error shuffling library: %s", err)
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

	ig := InputGame{}
	b, err := json.Marshal(game)
	if err != nil {
		log.Printf("failed to marshal input game in JoinGame: %s", err)
		return nil, err
	}
	err = json.Unmarshal(b, &ig)
	if err != nil {
		log.Printf("failed to unmarshal input game in JoinGame: %s", err)
		return nil, err
	}

	log.Printf("JoinGame#UpdateGame: %+v", ig)
	s.UpdateGame(ctx, ig)

	// s.mutex.Lock()
	// s.Directory[game.ID] = game
	// s.mutex.Unlock()
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
			log.Printf("error creating library from decklist: %+v", err)
			bs.Library = getCards(player.Library)
		} else {
			// Happy path
			bs.Library = library
		}

		commander, err := s.Card(ctx, player.Commander[0].Name, nil)
		if err != nil {
			log.Printf("error getting commander for deck: %+v", err)
			// fail gracefully and use their card name so they can still play a game
			inputCard := getCards(player.Commander)
			bs.Commander = []*Card{inputCard[0]}
		} else {
			bs.Commander = []*Card{commander[0]}
		}

		shuff, err := Shuffle(bs.Library)
		if err != nil {
			log.Printf("error shuffling library: %s", err)
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
	updated, err := boardStateFromInput(bs)
	if err != nil {
		log.Printf("UpdateBoardState failed to marshal input board state correctly: %s", err)
		return nil, fmt.Errorf("failed to marshal input boardstate: %s", err)
	}
	boardKey := BoardStateKey(bs.GameID, bs.User.Username)
	err = s.Set(boardKey, updated)
	if err != nil {
		log.Printf("error updating boardstate in redis: %s", err)
	}

	s.mutex.Lock()
	s.boardChannels[bs.User.Username] <- updated
	s.mutex.Unlock()
	return updated, nil
}

// Save saves a Game to the Persistence. TODO
func (g *Game) Save(ctx context.Context, db persistence.Persistence) (*Game, error) {
	return nil, errors.New("not impl")
}

// Save saves a BoardState to the Persistence interface. TODO
func (bs *BoardState) Save(ctx context.Context, db persistence.Persistence) (*BoardState, error) {
	return nil, errors.New("not impl")
}

func pushBoardStateUpdate(ctx context.Context, observers []Observer, input InputBoardState) {
	for _, obs := range observers {
		log.Printf("observers being notified: %+v\n", obs)
		log.Printf("board state updated: %+v\n", input)
	}
}

// gameFromInput transforms an InputGame to a *Game type
func gameFromInput(game InputGame) *Game {
	out := &Game{
		ID:        game.ID,
		PlayerIDs: getPlayerIDs(game.PlayerIDs),
	}
	if game.Turn == nil {
		out.Turn = &Turn{
			Player: game.Turn.Player,
			Phase:  game.Turn.Phase,
			Number: game.Turn.Number,
		}
	}

	if game.CreatedAt != nil {
		fmt.Printf("gameFromInput#createdAt: %+v\n", *game.CreatedAt)
		out.CreatedAt = *game.CreatedAt
	}

	if game.Handle != nil {
		out.Handle = game.Handle
	}

	return out
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

func boardStateFromInput(bs InputBoardState) (*BoardState, error) {
	data, err := json.Marshal(bs)
	if err != nil {
		return nil, errs.New("failed to marshal input game: %s", err)
	}
	new := &BoardState{}
	err = json.Unmarshal(data, &new)
	if err != nil {
		return nil, errs.New("failed to unmarshal game: %s", err)
	}

	return new, nil
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

// GameKey formats the keys for Games in our Directory
func GameKey(gameID string) string {
	return fmt.Sprintf("%s", gameID)
}

// TODO: Need to access Set and Get through Persistence interface instead
// of directly from the Server object. Persistence should be created and
// attached to the Server instead of directly attaching the Redis client.

// Set will set a value into the Redis client and returns an error, if any
func (s *graphQLServer) Set(key string, value interface{}) error {
	// TODO: Need to set this to an env variable
	exp, err := time.ParseDuration("12h")
	if err != nil {
		exp = 0
	}
	p, err := json.Marshal(value)
	if err != nil {
		log.Printf("failed to marshal value in Set: %s", err)
		return err
	}

	return s.redisClient.Set(key, p, exp).Err()
}

// Get returns a value from Redis client to `dest` and returns an error, if any
func (s *graphQLServer) Get(key string, dest interface{}) error {
	p, err := s.redisClient.Get(key).Result()
	if err != nil {
		log.Printf("failed to get key from redis: %s", err)
		return err
	}
	return json.Unmarshal([]byte(p), dest)
}
