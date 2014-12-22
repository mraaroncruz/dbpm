package commands

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

const _port = "4567"

// Server runs the web server that runs thepickmachine
func Server() {
	router := httprouter.New()
	router.GET("/", index)
	router.GET("/hello/:name", hello)

	port := os.Getenv("PORT")
	if port == "" {
		port = _port
	}

	fmt.Printf("Starting server on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}
