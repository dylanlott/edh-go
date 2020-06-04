package server

import (
	"context"
	"log"

	"github.com/zeebo/errs"
)

func (s *graphQLServer) Cards(
	ctx context.Context,
	id *string,
	name *string,
) ([]Card, error) {
	if name == nil {
		return nil, errs.New("must provide name for card")
	}

	rows, err := s.cardDB.Query(`SELECT "id", "name", "colors", "colorIdentity",
		"convertedManaCost", "manaCost", "uuid", "power", "toughness", "types",
		"subtypes", "supertypes", "isTextless", "text", "tcgplayerProductId"
		FROM "cards" WHERE "name" = ?`, name)
	if err != nil {
		return nil, errs.New("failed to run query: %s", err)
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
		data := make()
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

func (s *graphQLServer) Save(ctx context.Context, list []Card) ([]Card, error) {
	return nil, errs.New("not impl")
}
