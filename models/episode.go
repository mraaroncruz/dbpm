package models

import (
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
)

// Episode models an episode in the db
type Episode struct {
	ID, Number               int
	ShowID                   int       `db:"show_id"`
	PublishedAt              time.Time `db:"published_at"`
	Title, Slug, Description sql.NullString
}

// Save saves Episode struct to database
func (e *Episode) Save(db *sqlx.DB) error {
	query := `
    INSERT INTO episodes (
      number, published_at, title, slug, description)
    VALUES(
      :number, :published_at, :title, :slug, :description)`
	db.NamedExec(query, &e)
	return nil
}
