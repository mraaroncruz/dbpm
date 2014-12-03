package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/agonopol/readability"
)

type episode struct {
	Title, Slug, Description, Number string
	PublishedAt                      string `json:"published_at"`
	Picks                            []pick
}

type pick struct {
	Host, Name, Link, Description, Content string
}

func main() {
	url := "https://api.devchat.tv/show/ruby-rogues/episodes.json"
	client := http.Client{}
	res, err := client.Get(url)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var episodes = []episode{}
	err = json.Unmarshal(body, &episodes)
	if err != nil {
		panic(err)
	}

	pickChan := make(chan pick)
	pickCount := len(episodes[0].Picks)
	for _, aPick := range episodes[0].Picks {
		go func(p pick) {
			link := p.Link
			doc, err := readability.ParseURL(link)
			if err != nil {
				fmt.Printf("\n\nERRRROOORRRR: %s\n\n", err)
				pickChan <- p
				return
			}
			content, err := doc.Content()
			if err != nil {
				panic(err)
			}
			p.Content = content
			pickChan <- p
		}(aPick)
	}

	for i := 0; i < pickCount; i++ {
		currentPick := <-pickChan
		fmt.Printf("%#v\n", currentPick)
	}
}
