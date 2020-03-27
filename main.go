package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	h := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	/*
		# TODO:
		* start redis
		* connect to redis and init gamestate
		* register handlers to socket
		* listen on sockets
	*/

	http.HandleFunc("/", h)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
