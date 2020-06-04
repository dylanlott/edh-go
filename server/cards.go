package server

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/zeebo/errs"
)

func (s *graphQLServer) Cards(
	ctx context.Context,
	name string,
	id *string,
) ([]*Card, error) {
	if name == "" {
		return nil, errs.New("must provide name for card")
	}
	fmt.Printf("id: %+v \n name: %+v", id, name)
	rows, err := s.cardDB.Query(`SELECT "id", "name", "colors", "colorIdentity",
		"convertedManaCost", "manaCost", "uuid", "power", "toughness", "types",
		"subtypes", "supertypes", "isTextless", "text", "tcgplayerProductId"
		FROM "cards" WHERE "name" = ?`, name)
	if err != nil {
		return nil, errs.New("failed to run query: %s", err)
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
			Colors:        *colors,
			ColorIdentity: *colorIdentity,
			Cmc:           *convertedManaCost,
			ManaCost:      *manaCost,
			UUID:          *uuid,
			Power:         *power,
			Toughness:     *toughness,
			Types:         *types,
			Subtypes:      *subtypes,
			Supertypes:    *supertypes,
			IsTextless:    string(*isTextless),
			Text:          *text,
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

func (s *graphQLServer) Save(ctx context.Context, list []Card) ([]Card, error) {
	return nil, errs.New("not impl")
}
