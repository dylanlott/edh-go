package main

import (
	"fmt"
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
)

func main() {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}
	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		server.Emit("")
		fmt.Println("connected:", s.ID())
		return nil
	})
	server.OnError("/", func(s socketio.Conn, err error) {
		fmt.Println("ERRSOCKET:", err)
	})
	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("DISCONNECT:", reason)
	})
	go server.Serve()
	defer server.Close()

	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./web/dist/")))
	log.Println("Serving at localhost:9090...")
	log.Fatal(http.ListenAndServe(":9090", nil))
}
