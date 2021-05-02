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
		return nil, errors.New("not implemented")
	}
	g := &Game{}
	err := s.Get(GameKey(*gameID), g)
	if err != nil {
		log.Printf("games failed to get game %s: %s", *gameID, err)
		return nil, fmt.Errorf("failed to get game %s: %s", *gameID, err)
	}
	return []*Game{g}, nil
}

func (s *graphQLServer) GameUpdated(ctx context.Context, updated InputGame) (<-chan *Game, error) {
	// Update game in the directory
	b, err := json.Marshal(updated)
	if err != nil {
		return nil, errs.New("failed to marshal input game: %s", err)
	}
	game := &Game{}
	err = json.Unmarshal(b, &game)
	if err != nil {
		return nil, errs.New("failed to unmarshal game: %s", err)
	}

	s.mutex.Lock()
	// assign the game to the directory for finding
	// create a new gameChannel to announce Game updates over
	games := make(chan *Game, 1)
	// set the gameChannels to have the new receiving channel
	s.gameChannels[updated.ID] = games
	// announce the game over the GameChannels
	s.mutex.Unlock()
	games <- game
	go func() {
		<-ctx.Done()
		s.mutex.Lock()
		delete(s.gameChannels, updated.ID)
		s.mutex.Unlock()
	}()

	return games, nil
}

// UpdateGame is what's used to change the name of the game, format, insert
// or remove players, or change other meta informatin about a game.
// NB: Game _can_ touch boardstate right now, and it probably shouldn't.
func (s *graphQLServer) UpdateGame(ctx context.Context, new InputGame) (*Game, error) {

	// check existence of game, fail if not found
	b, err := json.Marshal(new)
	if err != nil {
		return nil, errs.New("failed to marshal input game: %s", err)
	}
	game := &Game{}
	err = json.Unmarshal(b, &game)
	if err != nil {
		return nil, errs.New("failed to unmarshal game: %s", err)
	}

	err = s.Set(GameKey(new.ID), game)
	if err != nil {
		return nil, fmt.Errorf("failed to save updated game state: %s", err)
	}

	s.mutex.Lock()
	s.gameChannels[new.ID] <- game
	s.mutex.Unlock()

	go s.GameUpdated(ctx, new)

	return game, nil
}

// JoinGame ...
func (s *graphQLServer) JoinGame(ctx context.Context, input *InputJoinGame) (*Game, error) {
	// TODO: Check context for User auth and append user info that way
	// TODO: We check for game existence a lot, we should probably make this a function
	if input.User.ID == nil {
		return nil, errors.New("must provide a user ID to join a game")
	}
	if input.User.Username == "" {
		return nil, errors.New("must provide a username to join a game")
	}

	game := &Game{}
	err := s.Get(GameKey(input.ID), &game)
	if err != nil {
		return nil, fmt.Errorf("failed to get game to join: %s", err)
	}

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
	err = s.Set(BoardStateKey(game.ID, user.Username), bs)
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

	log.Printf("unmarshaled input game: %+v", ig)
	go s.UpdateGame(ctx, ig)
	log.Printf("returning game: %+v", game)
	return game, nil
}

// createGame is untested currently
func (s *graphQLServer) CreateGame(ctx context.Context, inputGame InputCreateGame) (*Game, error) {
	// accept a game ID but create one if it isn't assigned
	if inputGame.ID == "" {
		inputGame.ID = uuid.New().String()
	}

	g := &Game{
		ID:        inputGame.ID,
		CreatedAt: time.Now(),
		PlayerIDs: []*User{},
		Turn: &Turn{
			Player: inputGame.Turn.Player,
			Phase:  inputGame.Turn.Phase,
			Number: inputGame.Turn.Number,
		},
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
			ID:       *player.User.ID,
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
	s.mutex.Unlock()

	// set *Game to Redis
	err := s.Set(GameKey(g.ID), g)
	if err != nil {
		return nil, fmt.Errorf("failed to save created game to redis: %s", err)
	}

	return g, nil
}

// func getPlayerIDs(inputUsers []*InputUser) []*User {
// 	var u []*User
// 	for _, i := range inputUsers {
// 		u = append(u, &User{
// 			ID:       *i.ID,
// 			Username: i.Username,
// 		})
// 	}

// 	return u
// }

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
