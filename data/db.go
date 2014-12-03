package data

import (
	"database/sql"
	"fmt"
)

type db struct {
	conn *sql.DB
}

// DB this is how you get ahold of the db from outside
var DB db

func init() {
	conn, err := sql.Open("postgres", "user=al dbname=dbpm_development port=4444 sslmode=verify-full")
	if err != nil {
		fmt.Errorf("DB connection error: %s", err)
	}
	DB = db{conn}
}
