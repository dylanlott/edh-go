package game

import (
	"fmt"
	"log"
	"math/rand"
	"strings"

	"github.com/dylanlott/edh-go/persistence"

	sdk "github.com/MagicTheGathering/mtg-sdk-go"
	"github.com/zeebo/errs"
)

// Card tracks the properties of a Card in a given Game
type Card struct {
	Name string

	// Track counters on a card
	Counters map[string]Counter

	// Data gets populated on query
	Data Data

	// wrappers around the mtg sdk card api
	CardInfo sdk.Card
	ID       sdk.CardId
}

// Data is used for populating card data with Query.
type Data map[string]interface{}

// CardList exposes a set of methods for manipulating a list of Cards
type CardList []Card

// Deck is the top level resource for a given Deck
type Deck struct {
	Name      string
	Commander CardList
	Format    string
	Cards     CardList
	Owner     UserID
}

// NewDecklist creates a new CardList from a line delimited list of card names.
// These names should be exact. This can be used for any format of Magic game,
// validation should be done in separate functions. This should purely be used
// to get the card's ID from MTG SDK ID.
func NewDecklist(db persistence.Database, raw string) (CardList, []error) {
	list := strings.Split(raw, "\n")
	decklist := make(CardList, 0, 99)
	errors := []error{}

	for _, i := range list {
		if i == "" {
			fmt.Println("continuing")
			continue
		}
		trimmed := strings.TrimSpace(i)
		card := Card{
			Name: trimmed,
		}

		card, err := getCard(db, trimmed)
		if err != nil {
			errors = append(errors, errs.Wrap(err))
			continue
		}

		log.Printf("retrieved card: %+v\n", card)

		decklist = append(decklist, card)
	}

	fmt.Printf("decklist: %+v\n", decklist)
	fmt.Printf("errors: %+v\n", errors)

	return decklist, errors
}

// Query will try to find card info for Card.Name
func Query(db persistence.Database, name string, id *string) (Card, error) {
	if name == "" {
		return Card{}, errs.New("must provide name for card")
	}

	rows, err := db.Query(`SELECT "id", "name", "colors", "colorIdentity",
		"convertedManaCost", "manaCost", "uuid", "power", "toughness", "types",
		"subtypes", "supertypes", "isTextless", "text", "tcgplayerProductId"
		FROM "cards" WHERE "name" = ?`, name)
	if err != nil {
		return Card{}, errs.New("failed to run query: %s", err)
	}

	cards := []Card{}

	for rows.Next() {
		var (
			id                 *int
			name               *string
			colors             *string
			colorIdentity      *string
			convertedManaCost  *string
			manaCost           *string
			uuid               *string
			power              *string
			toughness          *string
			types              *string
			subtypes           *string
			supertypes         *string
			isTextless         *int
			text               *string
			tcgplayerProductId *int
		)

		if err := rows.Scan(&id, &name, &colors, &colorIdentity,
			&convertedManaCost, &manaCost, &uuid, &power, &toughness, &types,
			&subtypes, &supertypes, &isTextless, &text,
			&tcgplayerProductId); err != nil {
			log.Printf("error scanning rows for card query: %s", err)
			continue
		}

		data := make(Data)
		data["name"] = *name
		data["id"] = *id

		card := Card{
			Name: *name,
			Data: data,
		}

		cards = append(cards, card)
	}
	// TODO: return card with given id if *id is passed to args

	return cards[0], err
}

// Shuffle is a sugar method to make Shuffling a list of Cards easier.
func (c CardList) Shuffle() (CardList, error) {
	deck := c
	rand.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})
	return deck, nil
}

// Validate will valiate the Deck against the format specified in args.
func (deck Deck) Validate(format string) bool {
	switch deck.Format {
	case "commander":
		break
	case "modern":
		break
	case "standard":
		break
	default:
		log.Printf("must provide deck format")
		return false
	}

	return false
}

// Fetch removes a card from the CardList and then shuffles the deck.
func Fetch(card Card, list CardList) (CardList, Card, error) {
	// TODO: Should we consider implementing opponent cuts here?
	found := false
	fetched := Card{}

	for _, c := range list {
		if card.Name == c.Name {
			found = true
			fmt.Printf("found card in decklist: %+v\n", card)
			fetched = c
			// TODO: Remove at index from slice
			break
		}
	}

	// OPINION: Anytime the player "touches" the deck, it should be shuffled.
	// That means there should be no path out where fetching doesn't shuffle
	// the deck, I whether the fetched card was found or not.
	if found == false {
		shuffled, err := list.Shuffle()
		if err != nil {
			return shuffled, Card{}, errs.New("failed to shuffle deck or find card")
		}
		return shuffled, Card{}, errs.New("card not in deck")
	}

	shuffled, err := list.Shuffle()
	if err != nil {
		return shuffled, Card{}, errs.New("failed to shuffled after successfully fetching")
	}

	return shuffled, fetched, nil
}

// Returns the top (0 indexed) card of the Deck into the player's Hand
// This means decks are drawn from left to right in an array, and the "bottom"
// of the deck is the last in the array.
func Draw(deck CardList, number int) (CardList, CardList) {
	// NB: if a player draws all of their cards, they don't lose. But if a player
	// would go to draw a card and there are none left, then they lose.
	if number > len(deck) {
		log.Printf("more cards were drawn than existed in deck")
		return nil, nil
	}
	cards := deck[number:]
	fmt.Printf("cards: %+v\n", cards)

	if len(deck) > 0 {
		deck = removeFromTop(deck, number)
	} else {
		// TODO: This technically means the player lost the game and we need to
		// account for that.
		log.Printf("deck was drawn when no cards were left")
		return nil, nil
	}

	return deck, cards
}

// remove from top will remove from the 0th position first and towards the right
func removeFromTop(deck CardList, i int) CardList {
	return append(deck[:i], deck[i+1:]...)
}

// removeAtIndex will remove an item from a slice at index `i` and return the
// updates slice.
func removeAtIndex(s CardList, index int) CardList {
	return append(s[:index], s[index+1:]...)
}

// getCard returns a single Card from the Database layer, or an error.
// If the card does not exist, an error will be thrown and Card{} will be
// returned. This is safe to run asynchronously.
func getCard(db persistence.Database, name string) (Card, error) {
	card, err := Query(db, name, nil)
	if err != nil {
		fmt.Printf("error querying for card in getCard: %+v\n", err)
		return Card{}, errs.Wrap(err)
	}
	return card, nil
}
