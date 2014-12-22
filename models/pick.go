package models

import (
	"bitbucket.org/pferdefleisch/dbpm/clients"
	"github.com/jmoiron/sqlx"
)

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
      :episode_id, :host, :name, :link, :description, :content)
		RETURNING id`
	rows, err := db.NamedQuery(query, &p)
	defer rows.Close()
	if err != nil {
		return err
	}

	if rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			return err
		}
		p.ID = id
	}
	return nil
}

// UpdateContent just updates the content field of pick
func (p *Pick) UpdateContent(db *sqlx.DB) error {
	query := "UPDATE picks SET content=:content WHERE picks.id = :id"
	db.NamedExec(query, &p)
	return nil
}

// ParseAPIPick takes the values from the api pick
// and adds them to the Pick
func (p *Pick) ParseAPIPick(apiPick *clients.Pick) error {
	p.Name = apiPick.Name
	p.Host = apiPick.Host
	p.Link = apiPick.Link
	p.Description = apiPick.Description
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
