package game

import (
	"log"
	"math/rand"
	"strings"

	"github.com/zeebo/errs"

	"github.com/dylanlott/edh-go/persistence"
)

// Card tracks the properties of a Card in a given Game. Cards can have their
// own state and their own interactions, since they're the atomic unit of
// Magic.
type Card struct {
	Name  string
	Split string // For the case where cards have two-in-one.

	// Data gets populated by database queries
	Data CardData

	// Game interactions
	CardTypes []string
	Counters  map[string]Counter
	Statuses  []string
	Blocking  CardList
}

// CardList exposes a set of methods for manipulating a list of Cards
type CardList []Card

type CardData map[string]interface{}

// Deck is the top level resource for a given Deck
type Deck struct {
	Name      string
	Commander CardList
	Format    string
	Cards     CardList
	Owner     UserID
}

// NewDecklist creates a new CardList from a line delimited list of card names.
// These names should be exact. This can be used for any format of Magic game.
// Validation should be done in separate functions. This function uses the
// SQLite database, so tests require it to be mocked.
func NewDecklist(db persistence.Database, raw string) (CardList, []error) {
	list := strings.Split(raw, "\n")
	decklist := make(CardList, 0, 99)
	errors := []error{}

	for _, i := range list {
		if i == "" {
			continue
		}
		trimmed := strings.TrimSpace(i)
		card, err := getCard(db, trimmed)
		if err != nil {
			errors = append(errors, errs.Wrap(err))
			continue
		}
		decklist = append(decklist, card)
	}

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

		// Add the data to a map for returning
		data := make(CardData)
		data["name"] = *name
		data["id"] = *id
		data["colors"] = *colors
		data["colorIdentity"] = *colorIdentity
		data["convertedManaCost"] = *convertedManaCost
		data["manaCost"] = *manaCost

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
func Shuffle(deck CardList) (CardList, error) {
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
func Fetch(card Card, list CardList) (Card, CardList, error) {
	// TODO: Should we consider implementing opponent cuts here?
	found := false
	fetched := Card{}

	for _, c := range list {
		if card.Name == c.Name {
			found = true
			fetched = c
			// TODO: Remove at index from slice
			break
		}
	}

	// OPINION: Anytime the player "touches" the deck, it should be shuffled.
	// That means there should be no path out where fetching doesn't shuffle
	// the deck, I whether the fetched card was found or not.
	if found == false {
		shuffled, err := Shuffle(list)
		if err != nil {
			return Card{}, shuffled, errs.New("failed to shuffle deck or find card")
		}
		return Card{}, shuffled, errs.New("card not in deck")
	}

	shuffled, err := Shuffle(list)
	if err != nil {
		return Card{}, shuffled, errs.New("failed to shuffled after successfully fetching")
	}

	return fetched, shuffled, nil
}

// Draw returns the top `number` cards of the deck, the shuffled library with
// the cards removed from it (drawn from it), and an error.
// Error will be thrown if a player tries to draw from an empty library, losing
// them the game.
func Draw(deck CardList, number int) (drawn CardList, library CardList, err error) {
	// NB: Drawing on an empty deck is a loss condition
	if len(deck) == 0 {
		return nil, nil, errs.New("check yourself before you deck yourself")
	}
	// NB: If a player draws all of their cards, they don't lose. But if a player
	// would go to draw a card and there are none left, then they lose.
	if number > len(deck) {
		return nil, nil, errs.New("check yourself before you deck yourself")
	}

	// draw the cards out
	drawn = deck[:number]
	library = deck[number:]

	return drawn, library, err
}

// Put inserts a CardList into a Deck at position `pos` in the deck.
// 0 is the top (the card that will be drawn next) and len(Deck)-1 is the
// bottom. The order of the cards inserted is preserved since this is a
// common library interaction.
func Put(deck CardList, cards CardList, pos int, shuffle bool) (library CardList, err error) {
	library = CardList{}
	fromTheTop := deck[:pos]
	bottom := deck[pos:]
	library = append(library, fromTheTop...)
	library = append(library, cards...)
	library = append(library, bottom...)

	if shuffle {
		shuffled, err := Shuffle(library)
		if err != nil {
			return CardList{}, errs.Wrap(err)
		}

		return shuffled, nil
	}

	return library, nil
}

// getCard returns a single Card from the Database layer, or an error.
// If the card does not exist, an error will be thrown and Card{} will be
// returned. This is safe to run asynchronously.
func getCard(db persistence.Database, name string) (Card, error) {
	card, err := Query(db, name, nil)
	if err != nil {
		return Card{}, errs.Wrap(err)
	}
	return card, nil
}
