package models

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"bitbucket.org/pferdefleisch/dbpm/clients"

	"github.com/jmoiron/sqlx"
)

// Episode models an episode in the db
type Episode struct {
	ID, Number               int
	ShowID                   int       `db:"show_id"`
	PublishedAt              time.Time `db:"published_at"`
	Title, Slug, Description sql.NullString
	Picks                    []clients.Pick `db:"-"`
}

// Save saves Episode struct to database
func (e *Episode) Save(db *sqlx.DB) error {
	query := `
    INSERT INTO episodes (
      show_id, number, published_at, title, slug, description)
    VALUES(
      :show_id, :number, :published_at, :title, :slug, :description)
		RETURNING id`
	rows, err := db.NamedQuery(query, &e)
	if err != nil {
		return err
	}
	defer rows.Close()

	var id int
	rows.Next()
	err = rows.Scan(&id)
	if err != nil {
		return err
	}
	e.ID = id
	return nil
}

// ParseAPIEpisode fill attributes from APIEpisode
func (e *Episode) ParseAPIEpisode(apiEpisode *clients.APIEpisode) {
	e.Number = apiEpisode.EpisodeNumber
	date, err := toDate(apiEpisode.PublishedAt)
	if err == nil {
		e.PublishedAt = date
	}
	e.Title = makeNullString(apiEpisode.TitleString)
	e.Slug = makeNullString(apiEpisode.Slug)
	e.Description = makeNullString(apiEpisode.Description)
	e.Picks = apiEpisode.Picks
}

func makeNullString(str string) sql.NullString {
	var null sql.NullString
	_ = null.Scan(str)
	return null
}

func toDate(dateString string) (time.Time, error) {
	if dateString != "" {
		devchatFormat := "01/02/06"
		theDate, err := time.Parse(devchatFormat, dateString)
		return theDate, err

	}
	return time.Time{}, errors.New("Date string from API empty")
}

// SavePicks takes persists the episode's picks to the database
func (e *Episode) SavePicks(db *sqlx.DB) ([]Pick, error) {
	modelPicks := []Pick{}
	for _, pick := range e.Picks {
		modelPick := &Pick{}
		modelPick.ParseAPIPick(&pick)
		modelPick.EpisodeID = e.ID
		err := modelPick.Save(db)
		if err != nil {
			fmt.Printf("Pick %s wouldn't save: %s\n", pick.Name, err)
		}
		modelPicks = append(modelPicks, *modelPick)
	}
	// never returns error
	return modelPicks, nil
}
