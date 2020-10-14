package server

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/zeebo/errs"
)

func (s *graphQLServer) Card(
	ctx context.Context,
	name string,
	id *string,
) ([]*Card, error) {
	if name == "" {
		return nil, errs.New("must provide name for card")
	}
	rows, err := s.cardDB.Query(`SELECT "id", "name", "colors", "colorIdentity",
		"convertedManaCost", "manaCost", "uuid", "power", "toughness", "types",
		"subtypes", "supertypes", "isTextless", "text", "tcgplayerProductId", 
		"scryfallIllustrationId" FROM "cards" WHERE "name" = ?`, name)
	if err != nil {
		return nil, errs.New("failed to run query: %s", err)
	}

	cards := []*Card{}
	for rows.Next() {
		var (
			id                     *int
			name                   *string
			colors                 *string
			colorIdentity          *string
			convertedManaCost      *string
			manaCost               *string
			uuid                   *string
			power                  *string
			toughness              *string
			types                  *string
			subtypes               *string
			supertypes             *string
			isTextless             *int
			text                   *string
			tcgplayerProductID     *int
			scryfallIllustrationID *string
		)

		if err := rows.Scan(&id, &name, &colors, &colorIdentity,
			&convertedManaCost, &manaCost, &uuid, &power, &toughness, &types,
			&subtypes, &supertypes, &isTextless, &text,
			&tcgplayerProductID, &scryfallIllustrationID); err != nil {
			log.Printf("error scanning rows for card query: %s", err)
			continue
		}

		parsedID := strconv.Itoa(*id)

		card := &Card{
			ID:            parsedID,
			Name:          *name,
			Colors:        colors,
			ColorIdentity: colorIdentity,
			Cmc:           convertedManaCost,
			ManaCost:      manaCost,
			UUID:          uuid,
			Power:         power,
			Toughness:     toughness,
			Types:         types,
			Subtypes:      subtypes,
			Supertypes:    supertypes,
			Text:          text,
			ScryfallID:    scryfallIllustrationID,
		}
		cards = append(cards, card)
	}
	// TODO: return card with given id if *id is passed to args

	if id != nil {
		for _, c := range cards {
			if c.ID == *id {
				return []*Card{c}, nil
			}
		}
	}

	return cards, err
}

func (s *graphQLServer) Cards(ctx context.Context, list []string) ([]*Card, error) {
	// TODO: Process `list` to allow for split cards
	query, args, err := sqlx.In(`SELECT "id", "name", "colors", "colorIdentity",
		"convertedManaCost", "manaCost", "uuid", "power", "toughness", "types",
		"subtypes", "supertypes", "isTextless", "text", "tcgplayerProductId" FROM cards WHERE name IN (?);`, list)
	if err != nil {
		return nil, errs.New("error formatting sqlx query")
	}

	rows, err := s.cardDB.Query(query, args...)
	if err != nil {
		return nil, errs.New("error querying cards DB for list of cards")
	}

	cards := []*Card{}

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

		parsedID := strconv.Itoa(*id)

		card := &Card{
			ID:            parsedID,
			Name:          *name,
			Colors:        colors,
			ColorIdentity: colorIdentity,
			Cmc:           convertedManaCost,
			ManaCost:      manaCost,
			UUID:          uuid,
			Power:         power,
			Toughness:     toughness,
			Types:         types,
			Subtypes:      subtypes,
			Supertypes:    supertypes,
			Text:          text,
		}
		cards = append(cards, card)
	}

	defer rows.Close()

	return cards, nil
}

// Search will search for card names in the database.
func (s *graphQLServer) Search(
	ctx context.Context,
	name *string,
	colors []*string,
	colorIdentity []*string,
	keywords []*string,
) ([]*Card, error) {
	if *name == "" {
		return nil, nil
	}
	n := fmt.Sprintf("%%%s%%", *name)
	rows, err := s.cardDB.Query("SELECT id, name, colors FROM cards WHERE name LIKE ?", n)
	if err != nil {
		log.Printf("error querying database: %s", err)
		return nil, errs.New("failed to search cardDB: %s", err)
	}

	cards := []*Card{}

	for rows.Next() {
		var (
			id     *int
			name   *string
			colors *string
		)

		if err := rows.Scan(&id, &name, &colors); err != nil {
			log.Printf("ERROR: failed to scan card into struct: %s", err)
			continue
		}

		parsedID := strconv.Itoa(*id)
		card := &Card{
			ID:     parsedID,
			Name:   *name,
			Colors: colors,
		}

		cards = append(cards, card)
	}

	return cards, nil
}

//
// Shuffle functions
//

// Shuffler type defines the interface for a given Shuffle function to fulfill.
type Shuffler func(deck []*Card) ([]*Card, error)

// Shuffle will apply a Knuth shuffle to the decklist.
func Shuffle(deck []*Card) ([]*Card, error) {
	rand.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})
	return deck, nil
}
