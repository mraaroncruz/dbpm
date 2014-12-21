package commands

import (
	"fmt"
	"log"

	"bitbucket.org/pferdefleisch/dbpm/data"
	"bitbucket.org/pferdefleisch/dbpm/models"
)

// Update checks the server for new episodes and parses their picks
func Update() {
	db := data.DBInstance()
	// for each shoUpdate(w
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

		apiEpisodes, err := clients.Devchat.GetEpisodesFrom(latestEpisodeNumber, show.Name)
		if err != nil {
			log.Fatalf("Couldn't get show episodes from api: %s\n", err)
		}
		//
		// for _, episode := range apiEpisodes {
		// 	err = episode.Save()
		// 	if err != nil {
		// 		fmt.Errorf("Couldn't save episode %s: %s\n", episode.Title, err)
		// 	}
		//
		// 	err = episode.SavePicks()
		// 	if err != nil {
		// 		fmt.Errorf("Couldn't save picks for %s: %s\n", episode.Title, err)
		// 	}
		// }
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
