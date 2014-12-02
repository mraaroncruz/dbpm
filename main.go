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
	Hosts, Guests                    []person
	PublishedAt                      string `json:"published_at"`
	Picks                            []pick
	Links                            link
}

type person struct {
	Name string
	Slug string
}

type pick struct {
	Host, Name, Link, Description string
}

type link struct {
	Episode, Show string
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
	for _, pick := range episodes[0].Picks {
		link := pick.Link
		doc, err := readability.ParseURL(link)
		if err != nil {
			panic(err)
		}
		content, err := doc.Content()
		fmt.Printf("%s\n", content)

	}
}
