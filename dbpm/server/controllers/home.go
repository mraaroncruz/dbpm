package controllers

import (
	"encoding/json"
	"net/http"

	"bitbucket.org/pferdefleisch/dbpm/dbpm/models"

	"github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"
)

// Home is the home controller
type Home struct {
	DB *sqlx.DB
}

// Index is the "landing page" for the app
func (c Home) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	db := c.DB
	pickModel := &models.Pick{}
    picks, err := pickModel.Latest(db)
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
