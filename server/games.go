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

	// Problem: There are multiple users at this point with the same UserIDs
	// log.Printf("there should not be multiple game.PlayerIDs: %+v\n", game.PlayerIDs)
	// This prints multiples, so where else are we setting boardstates?

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
	} else {
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
}

func getUsers(players []*InputUser) []*User {
	log.Printf("getUsers input: %+v\n", players)
	users := []*User{}
	for _, p := range players {
		u := &User{
			ID:       *p.ID,
			Username: p.Username,
		}
		users = append(users, u)
		log.Printf("users after append: %+v\n", users)
	}

	log.Printf("getUsers returning: %+v\n", users)
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

	log.Printf("#### game.PlayerIDs: %+v\n", game.PlayerIDs)

	output := &Game{
		ID:        found.ID,
		Handle:    game.Handle,
		CreatedAt: found.CreatedAt,
		PlayerIDs: getUsers(game.PlayerIDs),
		Turn:      getTurn(game.Turn),
		Rules:     found.Rules,
	}

	log.Printf("#GameUpdated#output: %+v\n", output)

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
	log.Printf("UpdateGame called with %+v\n", new)
	// check existence of game, fail if not found
	old, ok := s.Directory[new.ID]
	if !ok {
		return nil, errs.New("Game with ID %s does not exist", new.ID)
	}

	// update old game with new game data
	if err := mergo.Merge(&new, old); err != nil {
		return nil, errs.New("Failed to merge old game with new game: %s", err)
	}
	fmt.Printf("new game: %+v\n", &new)

	// cast new game into Game for GraphQL
	game := &Game{}
	if err := mergo.Merge(game, new); err != nil {
		return nil, errs.New("Failed to merge new game: %s", err)
	}

	fmt.Printf("fully updated game: %+v\n", game)

	s.mutex.Lock()
	s.Directory[new.ID] = game
	s.gameChannels[new.ID] <- game
	s.mutex.Unlock()

	return game, nil
}

// createGame is untested currently
func (s *graphQLServer) CreateGame(ctx context.Context, inputGame InputCreateGame) (*Game, error) {
	g := &Game{
		ID:        uuid.New().String(),
		CreatedAt: time.Now(),
		PlayerIDs: []*User{},
		// NB: Turns get added once the game has "started".
		// This is after roll for turn and mulligans happen.
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
	updated := boardStateFromInput(bs)
	fmt.Printf("UpdateBoardState hit: %+v\n", updated)
	boardKey := BoardStateKey(bs.GameID, bs.User.Username)
	err := s.Set(boardKey, updated)
	if err != nil {
		log.Printf("error updating boardstate in redis: %s", err)
	}

	s.mutex.Lock()
	log.Printf("pushing updated boardstate across channels: %+v", updated)
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

// gameFromInput transforms an InputGame to a *Game type
func gameFromInput(game InputGame) *Game {
	log.Printf("#gameFromInput#game: %+v\n", game)
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
		out.CreatedAt = *game.CreatedAt
	}

	if game.Handle != nil {
		out.Handle = game.Handle
	}

	return out
}

func getPlayerIDs(inputUsers []*InputUser) []*User {
	log.Printf("#getPlayerIDs#inputUsers: %+v\n", inputUsers)
	var u []*User
	for _, i := range inputUsers {
		u = append(u, &User{
			ID:       *i.ID,
			Username: i.Username,
		})
	}

	log.Printf("#getPlayerIDs#u<returning>: %+v\n", u)
	return u
}

func boardStateFromInput(bs InputBoardState) *BoardState {
	out := &BoardState{
		Life: bs.Life,
		User: &User{
			Username: bs.User.Username,
		},
		GameID: bs.GameID,
	}

	for _, c := range bs.Commander {
		out.Commander = append(out.Commander, &Card{
			Name:     c.Name,
			ID:       *c.ID,
			Quantity: c.Quantity,
			Tapped:   c.Tapped,
			Flipped:  c.Flipped,
			// TODO: Handle counters and labels
			// Counters:      c.Counters,
			Colors:        c.Colors,
			ColorIdentity: c.ColorIdentity,
			Cmc:           c.Cmc,
			ManaCost:      c.ManaCost,
			UUID:          c.UUID,
			Power:         c.Power,
			Toughness:     c.Toughness,
			Types:         c.Types,
			Subtypes:      c.Subtypes,
			Supertypes:    c.Supertypes,
			IsTextless:    c.IsTextless,
			Text:          c.Text,
			Tcgid:         c.Tcgid,
			ScryfallID:    c.ScryfallID,
		})
	}

	for _, c := range bs.Library {
		out.Library = append(out.Library, &Card{
			Name:     c.Name,
			ID:       *c.ID,
			Quantity: c.Quantity,
			Tapped:   c.Tapped,
			Flipped:  c.Flipped,
			// TODO: Handle counters and labels
			// Counters:      c.Counters,
			Colors:        c.Colors,
			ColorIdentity: c.ColorIdentity,
			Cmc:           c.Cmc,
			ManaCost:      c.ManaCost,
			UUID:          c.UUID,
			Power:         c.Power,
			Toughness:     c.Toughness,
			Types:         c.Types,
			Subtypes:      c.Subtypes,
			Supertypes:    c.Supertypes,
			IsTextless:    c.IsTextless,
			Text:          c.Text,
			Tcgid:         c.Tcgid,
			ScryfallID:    c.ScryfallID,
		})
	}

	for _, c := range bs.Exiled {
		out.Exiled = append(out.Exiled, &Card{
			Name:     c.Name,
			ID:       *c.ID,
			Quantity: c.Quantity,
			Tapped:   c.Tapped,
			Flipped:  c.Flipped,
			// TODO: Handle counters and labels
			// Counters:      c.Counters,
			Colors:        c.Colors,
			ColorIdentity: c.ColorIdentity,
			Cmc:           c.Cmc,
			ManaCost:      c.ManaCost,
			UUID:          c.UUID,
			Power:         c.Power,
			Toughness:     c.Toughness,
			Types:         c.Types,
			Subtypes:      c.Subtypes,
			Supertypes:    c.Supertypes,
			IsTextless:    c.IsTextless,
			Text:          c.Text,
			Tcgid:         c.Tcgid,
			ScryfallID:    c.ScryfallID,
		})
	}

	for _, c := range bs.Field {
		out.Field = append(out.Field, &Card{
			Name:     c.Name,
			ID:       *c.ID,
			Quantity: c.Quantity,
			Tapped:   c.Tapped,
			Flipped:  c.Flipped,
			// TODO: Handle counters and labels
			// Counters:      c.Counters,
			Colors:        c.Colors,
			ColorIdentity: c.ColorIdentity,
			Cmc:           c.Cmc,
			ManaCost:      c.ManaCost,
			UUID:          c.UUID,
			Power:         c.Power,
			Toughness:     c.Toughness,
			Types:         c.Types,
			Subtypes:      c.Subtypes,
			Supertypes:    c.Supertypes,
			IsTextless:    c.IsTextless,
			Text:          c.Text,
			Tcgid:         c.Tcgid,
			ScryfallID:    c.ScryfallID,
		})
	}

	for _, c := range bs.Hand {
		out.Hand = append(out.Hand, &Card{
			Name:     c.Name,
			ID:       *c.ID,
			Quantity: c.Quantity,
			Tapped:   c.Tapped,
			Flipped:  c.Flipped,
			// TODO: Handle counters and labels
			// Counters:      c.Counters,
			Colors:        c.Colors,
			ColorIdentity: c.ColorIdentity,
			Cmc:           c.Cmc,
			ManaCost:      c.ManaCost,
			UUID:          c.UUID,
			Power:         c.Power,
			Toughness:     c.Toughness,
			Types:         c.Types,
			Subtypes:      c.Subtypes,
			Supertypes:    c.Supertypes,
			IsTextless:    c.IsTextless,
			Text:          c.Text,
			Tcgid:         c.Tcgid,
			ScryfallID:    c.ScryfallID,
		})
	}

	for _, c := range bs.Controlled {
		out.Controlled = append(out.Controlled, &Card{
			Name:     c.Name,
			ID:       *c.ID,
			Quantity: c.Quantity,
			Tapped:   c.Tapped,
			Flipped:  c.Flipped,
			// TODO: Handle counters and labels
			// Counters:      c.Counters,
			Colors:        c.Colors,
			ColorIdentity: c.ColorIdentity,
			Cmc:           c.Cmc,
			ManaCost:      c.ManaCost,
			UUID:          c.UUID,
			Power:         c.Power,
			Toughness:     c.Toughness,
			Types:         c.Types,
			Subtypes:      c.Subtypes,
			Supertypes:    c.Supertypes,
			IsTextless:    c.IsTextless,
			Text:          c.Text,
			Tcgid:         c.Tcgid,
			ScryfallID:    c.ScryfallID,
		})
	}

	for _, c := range bs.Revealed {
		out.Revealed = append(out.Revealed, &Card{
			Name:     c.Name,
			ID:       *c.ID,
			Quantity: c.Quantity,
			Tapped:   c.Tapped,
			Flipped:  c.Flipped,
			// TODO: Handle counters and labels
			// Counters:      c.Counters,
			Colors:        c.Colors,
			ColorIdentity: c.ColorIdentity,
			Cmc:           c.Cmc,
			ManaCost:      c.ManaCost,
			UUID:          c.UUID,
			Power:         c.Power,
			Toughness:     c.Toughness,
			Types:         c.Types,
			Subtypes:      c.Subtypes,
			Supertypes:    c.Supertypes,
			IsTextless:    c.IsTextless,
			Text:          c.Text,
			Tcgid:         c.Tcgid,
			ScryfallID:    c.ScryfallID,
		})
	}

	return out
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

		quantity, err := strconv.ParseInt(record[0], 0, 64)
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
