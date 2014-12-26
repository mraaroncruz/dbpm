package clients

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

const (
	showsURL            = "https://api.devchat.tv/shows.json"
	parseableShowURL    = "https://api.devchat.tv/show/%s.json"
	parseableEpisodeURL = "https://api.devchat.tv/show/%s/episodes.json"
)

// APIEpisode is a representation of an episode in the api
type APIEpisode struct {
	Title, Slug, Number string
	Description         string
	TitleString         string `json:"title_string"`
	PublishedAt         string `json:"published_at"`
	Hosts               []person
	Guests              []person
	Picks               []Pick
	Links               link
	EpisodeNumber       int
}

func (e *APIEpisode) parse() error {
	var err error
	n := strings.TrimLeft(e.Number, "0")
	if n != "" {
		e.EpisodeNumber, err = strconv.Atoi(n)
	} else {
		err = errors.New("APIEpisode Number is blank")
		e.EpisodeNumber = 0
	}
	return err
}

type person struct {
	Name, Slug string
}

// Pick is a representation of a pick from the API
type Pick struct {
	Name, Host, Link, Description string
}

type link struct {
	Episode, Show string
}

type byNumber []APIEpisode

func (ben byNumber) Len() int           { return len(ben) }
func (ben byNumber) Swap(i, j int)      { ben[i], ben[j] = ben[j], ben[i] }
func (ben byNumber) Less(i, j int) bool { return ben[i].EpisodeNumber < ben[j].EpisodeNumber }

// Devchat is a client endpoint for the devchat.io API
type Devchat struct{}

// GetEpisodesAfter returns episodes after maxNum from the devchat api
func (devchat *Devchat) GetEpisodesAfter(maxNum int, showSlug string) (*[]APIEpisode, error) {
	episodes, err := devchat.GetEpisodes(showSlug)
	sort.Sort(byNumber(*episodes))
	newEpisodes := []APIEpisode{}
	for _, episode := range *episodes {
		if episode.EpisodeNumber > maxNum {
			newEpisodes = append(newEpisodes, episode)
		}
	}
	return &newEpisodes, err
}

// GetEpisodes returns episodes for a show by slug
func (devchat *Devchat) GetEpisodes(slug string) (*[]APIEpisode, error) {
	url := fmt.Sprintf(parseableEpisodeURL, slug)
	client := &http.Client{}
	res, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	rawData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	episodes := []APIEpisode{}
	err = json.Unmarshal(rawData, &episodes)
	if err != nil {
		return nil, err
	}

	for i := range episodes {
		epi := &episodes[i]
		err = epi.parse()
		if err != nil {
			fmt.Printf("%s\n", err)
		}
	}

	return &episodes, nil
}
