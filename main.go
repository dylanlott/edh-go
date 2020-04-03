package main

import (
	"fmt"
	"log"
	"net/http"

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
	_, err := persistence.NewRedis(config)
	if err != nil {
		log.Fatal(fmt.Errorf("error starting persistence: %s", err))
	}

	server, err := sockets.NewSocketLayer()
	if err != nil {
		log.Fatalf("failed to start socket server: %+v\n", err)
	}

	http.Handle("/socket.io/", server.GetClient())
	fmt.Printf("server listening on port 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
