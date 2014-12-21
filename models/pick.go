package models

import "github.com/jmoiron/sqlx"

// Pick model that is a db and json model
type Pick struct {
	ID          int
	EpisodeID   int `db:"episode_id"`
	Host        string
	Name        string
	Link        string
	Description string
	Content     string
}

// Save saves pick to database
func (p *Pick) Save(db *sqlx.DB) error {
	query := `
    INSERT INTO picks (
      episode_id, host, name, link, description, content)
    VALUES(
      :episode_id, :host, :name, :link, :description, :content)`
	db.NamedExec(query, &p)
	return nil
}

// PicksSearch searches for picks in db based on ++term++
func PicksSearch(term string, db *sqlx.DB) ([]Pick, error) {
	picks := []Pick{}
	query := `
    SELECT * FROM picks
      WHERE
        to_tsvector('english', name || ' ' || description || ' ' || content) @@
        to_tsquery('english', $1)`
	err := db.Select(&picks, query, term)
	if err != nil {
	}
	return picks, nil
}
