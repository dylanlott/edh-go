package main

import (
	"log"
	"net/http"
)

func main() {
	/*
		# TODO:
		* start redis
		* connect to redis and init gamestate
		* register handlers to socket
		* listen on sockets
	*/

	// serve web app
	srv := http.FileServer(http.Dir("./web"))
	log.Fatal(http.ListenAndServe(":6767", srv))
}
