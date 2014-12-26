package controllers

import (
	"encoding/json"
	"net/http"

	"bitbucket.org/pferdefleisch/dbpm/dbpm/models"

	"github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"
)

// Search controller
type Search struct {
	DB *sqlx.DB
}

// Index takes a query string `q` and a possible show and gives pick search results
// for that query
func (c Search) Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	db := c.DB
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

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsn)
}
