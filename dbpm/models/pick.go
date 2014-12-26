package models

import (
	"time"

	"bitbucket.org/pferdefleisch/dbpm/dbpm/clients"

	"github.com/jmoiron/sqlx"
)

// Pick model that is a db and json model
type Pick struct {
	ID          int    `db:"id"`
	EpisodeID   int    `db:"episode_id"`
	Host        string `db:"host"`
	Name        string `db:"name"`
	Link        string `db:"link"`
	Description string `db:"description"`
	Content     string `db:"content"`
}

// FullPick represents a pick with episode and show data
type FullPick struct {
	Host               string    `json:"host" db:"host"`
	Name               string    `json:"name" db:"name"`
	Link               string    `json:"link" db:"link"`
	Description        string    `json:"description" db:"description"`
	Number             string    `json:"number" db:"number"`
	EpisodeTitle       string    `json:"episode_title" db:"episode_title"`
	EpisodeDescription string    `json:"episode_description" db:"episode_description"`
	EpisodeSlug        string    `json:"episode_slug" db:"episode_slug"`
	ShowName           string    `json:"show_name" db:"show_name"`
	ShowSlug           string    `json:"show_slug" db:"show_slug"`
	PublishedAt        time.Time `json:"published_at" db:"published_at"`
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

// Latest gets a list of latest picks
func (p *Pick) Latest(db *sqlx.DB) ([]FullPick, error) {
	query := `
		SELECT
			p.host, p.name, p.link, p.description,
			e.number, e.published_at, e.title episode_title,
			e.slug episode_slug, e.description episode_description,
			s.name show_name, s.slug show_slug
		FROM picks AS p
	  LEFT JOIN episodes AS e
			ON p.episode_id = e.id
		LEFT JOIN shows AS s
			ON e.show_id = s.id
		ORDER BY e.published_at DESC
		LIMIT 30
	`
	picks := []FullPick{}
	err := db.Select(&picks, query)
	if err != nil {
		return nil, err
	}
	return picks, nil
}

// ShowSearch searches a show for picks
func (p *Pick) ShowSearch(db *sqlx.DB, term, showSlug string) ([]FullPick, error) {
	query := `
    SELECT
			p.host, p.name, p.link, p.description,
			e.number, e.published_at, e.title episode_title,
			e.slug episode_slug, e.description episode_description,
			s.name show_name, s.slug show_slug
		FROM picks AS p
	  LEFT JOIN episodes AS e
			ON p.episode_id = e.id
		LEFT JOIN shows AS s
			ON e.show_id = s.id
    WHERE
      to_tsvector('english', p.name || ' ' || p.description || ' ' || p.content) @@
      to_tsquery('english', $1)
		AND
		   s.slug = $2
		`
	picks := []FullPick{}
	err := db.Select(&picks, query, term, showSlug)
	if err != nil {
		return nil, err
	}
	return picks, nil
}

// AllSearch searches all shows for picks
func (p *Pick) AllSearch(db *sqlx.DB, term string) ([]FullPick, error) {
	query := `
    SELECT
			p.host, p.name, p.link, p.description,
			e.number, e.published_at, e.title episode_title,
			e.slug episode_slug, e.description episode_description,
			s.name show_name, s.slug show_slug
		FROM picks AS p
	  LEFT JOIN episodes AS e
			ON p.episode_id = e.id
		LEFT JOIN shows AS s
			ON e.show_id = s.id
    WHERE
      to_tsvector('english', p.name || ' ' || p.description || ' ' || p.content) @@
      to_tsquery('english', $1)
		`
	picks := []FullPick{}
	err := db.Select(&picks, query, term)
	if err != nil {
		return nil, err
	}
	return picks, nil
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
