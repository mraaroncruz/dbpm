package models

import (
	"time"

	"github.com/jmoiron/sqlx"
)

// Episode models an episode in the db
type Episode struct {
	ID, Number               int
	PublishedAt              time.Time `db:"published_at"`
	Title, Slug, Description string
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
