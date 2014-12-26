package models

import "github.com/jmoiron/sqlx"

// Show models a show like ruby rogues in the db
type Show struct {
	ID         int
	Name, Slug string
}

// ShowAll get all shows from database
func ShowAll(db *sqlx.DB) (*[]Show, error) {
	shows := []Show{}
	query := "SELECT * FROM shows"
	err := db.Select(&shows, query)
	return &shows, err
}

// ShowFind finds a show in the database by key
func ShowFind(key string) *Show {
	return new(Show)
}

// MaxEpisodeNumber returns the most recent episode from the
// list of episodes sorted by episode number
func (show *Show) MaxEpisodeNumber(db *sqlx.DB) (int, error) {
	episode, err := latestEpisode(show, db)
	if err != nil {
		return 0, err
	}
	return episode.Number, nil
}

func latestEpisode(show *Show, db *sqlx.DB) (*Episode, error) {
	var episodes = []Episode{}
	query := "SELECT * FROM episodes WHERE episodes.show_id = $1 ORDER BY number DESC LIMIT 1"
	err := db.Select(&episodes, query, show.ID)
	if len(episodes) > 0 {
		return &episodes[0], err
	}
	return new(Episode), err
}
