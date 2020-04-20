package api

import (
	"github.com/graphql-go/graphql"

	"github.com/dylanlott/edh-go/persistence"
)

// Server holds a reference to the graphql server and the Persistence server.
type Server struct {
	Database persistence.Persistence
}

// Game holds a reference to a game.
type Game struct {
	ID string
}

// User represents a player in the game and is tied to one account.
type User struct {
	ID string
}

// Message represents a user - to - user message, not a pub/sub message.
type Message struct {
	ID   string
	Body string
	User string
}

var gameType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Game",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"users": &graphql.Field{
				Type: graphql.NewList(userType),
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					// TODO: Fill this out
					return nil, nil
				},
			},
		},
	},
)

var userType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var messageType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Message",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
