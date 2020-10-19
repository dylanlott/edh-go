// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package server

import (
	"time"
)

type BoardState struct {
	User       *User      `json:"User"`
	Life       int        `json:"Life"`
	GameID     string     `json:"GameID"`
	Commander  []*Card    `json:"Commander"`
	Library    []*Card    `json:"Library"`
	Graveyard  []*Card    `json:"Graveyard"`
	Exiled     []*Card    `json:"Exiled"`
	Field      []*Card    `json:"Field"`
	Hand       []*Card    `json:"Hand"`
	Revealed   []*Card    `json:"Revealed"`
	Controlled []*Card    `json:"Controlled"`
	Counters   []*Counter `json:"Counters"`
}

type Card struct {
	Name          string     `json:"Name"`
	ID            string     `json:"ID"`
	Quantity      *int       `json:"Quantity"`
	Tapped        *bool      `json:"Tapped"`
	Flipped       *bool      `json:"Flipped"`
	Counters      []*Counter `json:"Counters"`
	Colors        *string    `json:"Colors"`
	ColorIdentity *string    `json:"ColorIdentity"`
	Cmc           *string    `json:"CMC"`
	ManaCost      *string    `json:"ManaCost"`
	UUID          *string    `json:"UUID"`
	Power         *string    `json:"Power"`
	Toughness     *string    `json:"Toughness"`
	Types         *string    `json:"Types"`
	Subtypes      *string    `json:"Subtypes"`
	Supertypes    *string    `json:"Supertypes"`
	IsTextless    *string    `json:"IsTextless"`
	Text          *string    `json:"Text"`
	Tcgid         *string    `json:"TCGID"`
	ScryfallID    *string    `json:"ScryfallID"`
}

type Counter struct {
	Name  string `json:"Name"`
	Value string `json:"Value"`
}

type Deck struct {
	ID        string   `json:"ID"`
	Name      string   `json:"Name"`
	Commander string   `json:"Commander"`
	Library   []string `json:"Library"`
}

type Emblem struct {
	Name   string `json:"Name"`
	Value  string `json:"Value"`
	Player *User  `json:"Player"`
}

type Game struct {
	ID        string    `json:"ID"`
	Handle    *string   `json:"Handle"`
	CreatedAt time.Time `json:"CreatedAt"`
	Rules     []*Rule   `json:"Rules"`
	Turn      *Turn     `json:"Turn"`
	PlayerIDs []*User   `json:"PlayerIDs"`
}

type InputBoardState struct {
	User       *InputUser      `json:"User"`
	GameID     string          `json:"GameID"`
	Life       int             `json:"Life"`
	Decklist   *string         `json:"Decklist"`
	Commander  []*InputCard    `json:"Commander"`
	Library    []*InputCard    `json:"Library"`
	Graveyard  []*InputCard    `json:"Graveyard"`
	Exiled     []*InputCard    `json:"Exiled"`
	Field      []*InputCard    `json:"Field"`
	Hand       []*InputCard    `json:"Hand"`
	Revealed   []*InputCard    `json:"Revealed"`
	Controlled []*InputCard    `json:"Controlled"`
	Counters   []*InputCounter `json:"Counters"`
	Emblems    []*InputEmblem  `json:"Emblems"`
}

type InputCard struct {
	ID            *string         `json:"ID"`
	Name          string          `json:"Name"`
	Counters      []*InputCounter `json:"Counters"`
	Labels        []*InputLabel   `json:"Labels"`
	Tapped        *bool           `json:"Tapped"`
	Flipped       *bool           `json:"Flipped"`
	Quantity      *int            `json:"Quantity"`
	Colors        *string         `json:"Colors"`
	ColorIdentity *string         `json:"ColorIdentity"`
	Cmc           *string         `json:"CMC"`
	ManaCost      *string         `json:"ManaCost"`
	UUID          *string         `json:"UUID"`
	Power         *string         `json:"Power"`
	Toughness     *string         `json:"Toughness"`
	Types         *string         `json:"Types"`
	Subtypes      *string         `json:"Subtypes"`
	Supertypes    *string         `json:"Supertypes"`
	IsTextless    *string         `json:"IsTextless"`
	Text          *string         `json:"Text"`
	Tcgid         *string         `json:"TCGID"`
	ScryfallID    *string         `json:"ScryfallID"`
}

type InputCounter struct {
	Card  *InputCard `json:"Card"`
	Name  string     `json:"Name"`
	Value string     `json:"Value"`
}

type InputCreateGame struct {
	ID      string             `json:"ID"`
	Turn    *InputTurn         `json:"Turn"`
	Handle  *string            `json:"Handle"`
	Players []*InputBoardState `json:"Players"`
}

type InputDeck struct {
	Name      *string  `json:"Name"`
	Commander []string `json:"Commander"`
	Cards     []string `json:"Cards"`
}

type InputEmblem struct {
	Name   string     `json:"Name"`
	Value  string     `json:"Value"`
	Player *InputUser `json:"Player"`
}

type InputGame struct {
	ID        string       `json:"ID"`
	Turn      *InputTurn   `json:"Turn"`
	CreatedAt *time.Time   `json:"Created_At"`
	Handle    *string      `json:"Handle"`
	PlayerIDs []*InputUser `json:"PlayerIDs"`
}

type InputLabel struct {
	Name       string `json:"Name"`
	Value      string `json:"Value"`
	AssignedBy string `json:"AssignedBy"`
}

type InputSignup struct {
	Username string `json:"Username"`
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

type InputTurn struct {
	Player string `json:"Player"`
	Phase  string `json:"Phase"`
	Number int    `json:"Number"`
}

type InputUser struct {
	Username string  `json:"Username"`
	ID       *string `json:"ID"`
}

type Message struct {
	ID        string    `json:"ID"`
	User      string    `json:"User"`
	CreatedAt time.Time `json:"CreatedAt"`
	Text      string    `json:"Text"`
	GameID    string    `json:"GameID"`
	Channel   *string   `json:"Channel"`
}

type Rule struct {
	Name  string `json:"Name"`
	Value string `json:"Value"`
}

type Turn struct {
	Player string `json:"Player"`
	Phase  string `json:"Phase"`
	Number int    `json:"Number"`
}

type User struct {
	ID       string `json:"ID"`
	Username string `json:"Username"`
	Deck     string `json:"Deck"`
}
