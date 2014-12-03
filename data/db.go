package data

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // import postgres driver
)

// DB this is how you get ahold of the db from outside
var DB *sql.DB

func init() {
	conn, err := sql.Open("postgres", "user=al dbname=dbpm_development port=4444 sslmode=disable")
	if err != nil {
		fmt.Errorf("DB connection error: %s", err)
	}
	DB = conn
}
