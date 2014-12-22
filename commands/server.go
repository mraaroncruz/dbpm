package commands

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"bitbucket.org/pferdefleisch/dbpm/data"
	"bitbucket.org/pferdefleisch/dbpm/server/controllers"

	"github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"
)

const _port = "4567"

var db *sqlx.DB

// Server runs the web server that runs thepickmachine
func Server() {
	db = data.DBInstance()
	router := httprouter.New()
	home := &controllers.Home{DB: db}
	router.GET("/", home.Index)
	search := &controllers.Search{DB: db}
	router.GET("/search", search.Index)

	port := os.Getenv("PORT")
	if port == "" {
		port = _port
	}

	fmt.Printf("Starting server on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
