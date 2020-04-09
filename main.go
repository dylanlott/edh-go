package main

import (
	"log"
	"net/http"

	_ "github.com/dylanlott/edh-go/persistence"
	"github.com/dylanlott/edh-go/sockets"
	"github.com/rs/cors"
)

// Old
// // // // /// // // // // // // // // // // // /// // // // // // // //
// func main() {
// 	/*
// 		# TODO:
// 		* [x] start redis
// 		* [x] connect to redis
// 		* [] listen on sockets
// 		* [] init gamestate
// 		* [] register handlers to socket
// 	*/
//
// 	config := make(persistence.Config)
// 	_, err := persistence.NewRedis(config)
// 	if err != nil {
// 		log.Fatal(fmt.Errorf("error starting persistence: %s", err))
// 	}
//
// 	server, err := sockets.NewSocketLayer()
// 	if err != nil {
// 		log.Fatalf("failed to start socket server: %+v\n", err)
// 	}
//
// 	c := cors.New(cors.Options{
//     AllowedOrigins:   []string{"*"},
//     AllowCredentials: true,
//   })
//
//   // decorate existing handler with cors functionality set in c
//   handler = c.Handler(handler)
// 	handler.
//
// 	http.Handle("/socket.io/", server.GetClient())
// 	fmt.Printf("server listening on port 8000")
// 	log.Fatal(http.ListenAndServe(":8000", nil))
// }

func main() {
	mux := http.NewServeMux()
	// mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.Write([]byte("{\"hello\": \"world\"}"))
	// })

	// cors.Default() setup the middleware with default options being
	// all origins accepted with simple methods (GET, POST). See
	// documentation below for more options.
	server, err := sockets.NewSocketLayer()
	if err != nil {
		log.Fatal(err)
	}

	mux.Handle("/socket.io/", server.GetClient())
	mux.Handle("/", http.FileServer(http.Dir("web/dist/")))

	// provide default cors to the mux
	handler := cors.Default().Handler(mux)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	// decorate existing handler with cors functionality set in c
	handler = c.Handler(handler)

	log.Println("Serving at localhost:8000...")
	log.Fatal(http.ListenAndServe(":8000", handler))
}
