package models

import "github.com/pferdefleisch/dbpm/data"

// Pick model that is a db and json model
type Pick struct {
	ID          int
	EpisodeID   int `database:"episode_id"`
	Host        string
	Name        string
	Link        string
	Description string
	Content     string
}

// Save saves pick to database
func (p *Pick) Save() error {
	db := data.DB
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(`
    INSERT INTO picks
      (episode_id, host, name, link, description, content)
    VALUES
      ($1, $2, $3, $4, $5, $6)`,
		p.EpisodeID,
		p.Host,
		p.Name,
		p.Link,
		p.Description,
		p.Content)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

// Search searches for picks in db based on ++term++
func Search(term string) ([]Pick, error) {
	return nil, nil
}
