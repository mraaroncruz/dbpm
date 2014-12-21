package commands

import (
	"fmt"
	"log"

	"bitbucket.org/pferdefleisch/dbpm/clients"
	"bitbucket.org/pferdefleisch/dbpm/data"
	"bitbucket.org/pferdefleisch/dbpm/models"
)

// Update checks the server for new episodes and parses their picks
func Update() {
	db := data.DBInstance()

	shows, err := models.ShowAll(db)
	if err != nil {
		log.Fatalf("Couldn't retrieve all songs.\n")
	}

	for _, show := range *shows {
		latestEpisodeNumber, err := show.MaxEpisodeNumber(db)
		if err != nil {
			log.Fatalf("Couldn't retreive latest episode from %s: %s\n", show.Name, err)
		}
		fmt.Printf("Latest: %s %d\n", show.Name, latestEpisodeNumber)

		var apiEpisodes = &[]clients.APIEpisode{}
		devchat := &clients.Devchat{}
		apiEpisodes, err = devchat.GetEpisodesAfter(latestEpisodeNumber, show.Slug)
		if err != nil {
			log.Fatalf("Couldn't get show episodes from api: %s\n", err)
		}
		fmt.Printf("%d new episodes\n", len(*apiEpisodes))

		for _, episode := range *apiEpisodes {
			dbEpisode := &models.Episode{}
			dbEpisode.ParseAPIEpisode(&episode)
			dbEpisode.ShowID = show.ID
			err = dbEpisode.Save(db)
			if err != nil {
				fmt.Printf("Couldn't save episode %s: %s\n", dbEpisode.Title, err)
			}

			err = dbEpisode.SavePicks(db)
			if err != nil {
				fmt.Printf("Couldn't save picks for %s: %s\n", episode.Title, err)
			}

			// err = episode.ScrapePicks()
			// if err != nil {
			// 	fmt.Printf("Error scraping picks for episode %s\n%s\n", episode.Title, err)
			// }
		}
		// fmt.Printf("Saved picks from %d episodes of %s\n", len(apiEpisodes), show.Name)
	}
	//   parse url
	//   get difference of episodes
	//   create new episode
	//   create each pick from episode
	//
	// db := data.DBInstance()
	// defer db.Close()

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
