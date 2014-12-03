package data

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func DBInstance() *sqlx.DB {
	db, err := sqlx.Connect("postgres", "user=al dbname=dbpm_development port=4444 sslmode=disable")
	if err != nil {
		fmt.Printf("DB connection error: %s\n\n", err)
	}
	return db
}
