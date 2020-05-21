// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package server

import (
	"time"
)

type BoardState struct {
	User       *User      `json:"User"`
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
	Name string `json:"Name"`
	ID   string `json:"ID"`
}

type Counter struct {
	Card  *Card  `json:"card"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Deck struct {
	ID        string   `json:"ID"`
	Name      string   `json:"Name"`
	Commander string   `json:"Commander"`
	Library   []string `json:"Library"`
}

type Emblem struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Player *User  `json:"player"`
}

type Game struct {
	ID        string        `json:"id"`
	CreatedAt time.Time     `json:"created_at"`
	Rules     []*Rule       `json:"rules"`
	Turn      *Turn         `json:"turn"`
	Players   []*BoardState `json:"players"`
}

type InputBoardState struct {
	UserID     string          `json:"UserID"`
	GameID     string          `json:"GameID"`
	Commander  []string        `json:"Commander"`
	Library    []*string       `json:"Library"`
	Graveyard  []*string       `json:"Graveyard"`
	Exiled     []*string       `json:"Exiled"`
	Field      []*string       `json:"Field"`
	Hand       []*string       `json:"Hand"`
	Revealed   []*string       `json:"Revealed"`
	Controlled []*string       `json:"Controlled"`
	Counters   []*InputCounter `json:"Counters"`
	Emblems    []*InputEmblem  `json:"Emblems"`
}

type InputCard struct {
	ID   *string `json:"ID"`
	Name string  `json:"Name"`
}

type InputCounter struct {
	Card  *InputCard `json:"card"`
	Name  string     `json:"name"`
	Value string     `json:"value"`
}

type InputDeck struct {
	Name      *string  `json:"name"`
	Commander []string `json:"commander"`
	Cards     []string `json:"cards"`
}

type InputEmblem struct {
	Name   string     `json:"name"`
	Value  string     `json:"value"`
	Player *InputUser `json:"player"`
}

type InputGame struct {
	Players []*InputUser `json:"players"`
}

type InputSignup struct {
	Username string `json:"Username"`
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

type InputUser struct {
	Deck     *InputDeck `json:"Deck"`
	Username string     `json:"Username"`
}

type Message struct {
	ID        string    `json:"id"`
	User      string    `json:"user"`
	CreatedAt time.Time `json:"createdAt"`
	Text      string    `json:"text"`
}

type Rule struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Turn struct {
	Player *string `json:"Player"`
	Phase  *string `json:"Phase"`
	Number *int    `json:"Number"`
}

type User struct {
	ID         string      `json:"id"`
	Username   string      `json:"username"`
	Deck       string      `json:"deck"`
	Boardstate *BoardState `json:"boardstate"`
}
