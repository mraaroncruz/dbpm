package commands

import (
	"fmt"
	"os"

	_ "github.com/lib/pq" // blah blah
	"github.com/pferdefleisch/dbpm/data"
	"github.com/pferdefleisch/dbpm/models"
)

// Search is a one-off search of the pick database
func Search(term string) {
	db := data.DBInstance()
	defer db.Close()
	picks, err := models.PicksSearch(term, db)
	if err != nil {
		fmt.Printf("Search error: %s\n", err)
		os.Exit(1)
	}
	for i, pick := range picks {
		fmt.Printf("%d: %#v\n\n", i+1, pick)
	}
}
