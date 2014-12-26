package commands

import (
	"fmt"
	"os"

	"bitbucket.org/pferdefleisch/dbpm/dbpm/data"
	"bitbucket.org/pferdefleisch/dbpm/dbpm/models"

	_ "github.com/lib/pq" // blah blah
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
