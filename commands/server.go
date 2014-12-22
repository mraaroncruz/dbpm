package commands

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"bitbucket.org/pferdefleisch/dbpm/data"
	"bitbucket.org/pferdefleisch/dbpm/models"

	"github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"
)

const _port = "4567"

var db *sqlx.DB

// Server runs the web server that runs thepickmachine
func Server() {
	db = data.DBInstance()
	router := httprouter.New()
	router.GET("/", index)
	router.GET("/search", search)

	port := os.Getenv("PORT")
	if port == "" {
		port = _port
	}

	fmt.Printf("Starting server on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	pickModel := &models.Pick{}
	picks, err := pickModel.Latest(db)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	jsn, err := json.Marshal(picks)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsn)
}

func search(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	term := r.FormValue("q")
	showSlug := r.FormValue("show")

	var err error
	picks := []models.FullPick{}
	pickModel := &models.Pick{}
	if showSlug == "" {
		picks, err = pickModel.AllSearch(db, term)
	} else {
		picks, err = pickModel.ShowSearch(db, term, showSlug)
	}

	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	jsn, err := json.Marshal(picks)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsn)
}
