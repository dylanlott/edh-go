package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/zeebo/errs"

	"github.com/dylanlott/edh-go/persistence"
	"github.com/dylanlott/edh-go/sockets"
)

func main() {
	/*
		# TODO:
		* [x] start redis
		* [x] connect to redis
		* [] listen on sockets
		* [] init gamestate
		* [] register handlers to socket
	*/

	config := make(persistence.Config)
	db, err := persistence.NewRedis(config)
	if err != nil {
		log.Fatal(fmt.Errorf("error starting persistence: %s", err))
	}

	fmt.Printf("connected to db - [%+v]\n", db)

	// start socketio
	s, err := sockets.NewFullSocketer()
	if err != nil {
		log.Fatal(errs.New("failed to start socket server"))
	}

	fmt.Printf("started socket server: %+v\n", s)

	// serve web app
	srv := http.FileServer(http.Dir("./web"))
	log.Fatal(http.ListenAndServe(":6767", srv))
}
