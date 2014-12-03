package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pferdefleisch/dbpm/models"
)

func createDB() *sqlx.DB {
	db, err := sqlx.Connect("postgres", "user=al dbname=dbpm_development port=4444 sslmode=disable")
	if err != nil {
		fmt.Printf("DB connection error: %s\n\n", err)
	}
	return db
}

type episode struct {
	Title, Slug, Description, Number string
	PublishedAt                      string `json:"published_at"`
	Picks                            []models.Pick
}

func main() {
	db := createDB()
	defer db.Close()
	picks, err := models.SearchPicks("amazon", db)
	if err != nil {
		fmt.Printf("Error is %s\n", err)
	}
	fmt.Printf("%#v\n", picks)

	// url := "https://api.devchat.tv/show/ruby-rogues/episodes.json"
	// client := http.Client{}
	// res, err := client.Get(url)
	// if err != nil {
	// 	fmt.Printf("ERRRRR::: %s\n\n\n", err)
	// }
	//
	// defer res.Body.Close()
	//
	// body, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Printf("ERRRRR::: %s\n\n\n", err)
	// }
	//
	// var episodes = []episode{}
	// err = json.Unmarshal(body, &episodes)
	// if err != nil {
	// 	fmt.Printf("ERRRRR::: %s\n\n\n", err)
	// }
	//
	// pickChan := make(chan models.Pick)
	// pickCount := len(episodes[0].Picks)
	// for _, aPick := range episodes[0].Picks {
	// 	go func(p models.Pick) {
	// 		link := p.Link
	// 		doc, err := readability.ParseURL(link)
	// 		if err != nil {
	// 			fmt.Printf("\n\nERRRROOORRRR: %s\n\n", err)
	// 			pickChan <- p
	// 			return
	// 		}
	// 		content, err := doc.Content()
	// 		if err != nil {
	// 			fmt.Printf("ERRRRR::: %s\n\n\n", err)
	// 		}
	// 		p.Content = content
	// 		pickChan <- p
	// 	}(aPick)
	// }
	//
	// for i := 0; i < pickCount; i++ {
	// 	currentPick := <-pickChan
	// 	err = currentPick.Save(db)
	// 	if err != nil {
	// 		fmt.Printf("ERRRRR::: %s\n\n\n", err)
	// 	}
	// 	fmt.Printf("%#v\n", currentPick)
	// }
}
