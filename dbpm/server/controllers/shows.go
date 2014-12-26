package controllers

import (
	"encoding/json"
	"net/http"

	"bitbucket.org/pferdefleisch/dbpm/dbpm/models"

	"github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"
)

// Shows is the shows controller
type Shows struct {
	DB *sqlx.DB
}

// Index gives a list of all the shows
func (c Shows) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	db := c.DB
	showModel := &models.Show{}
	shows, err := showModel.All(db)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	jsn, err := json.Marshal(shows)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsn)
}
