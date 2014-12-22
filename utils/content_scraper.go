package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/agonopol/readability"
	"github.com/jmoiron/sqlx"

	"bitbucket.org/pferdefleisch/dbpm/models"
)

// ContentScraper is a utility to concurrently scrape the content from picks
type ContentScraper struct {
	DB *sqlx.DB
}

// Scrape concurrently scrapes content from pick links
func (cs *ContentScraper) Scrape(picks []models.Pick) error {
	wg := sync.WaitGroup{}
	for i := 0; i < len(picks); i++ {
		wg.Add(1)
		scrapeContent(&wg, cs.DB, &(picks[i]))
	}
	wg.Wait()
	return nil
}

func scrapeContent(wg *sync.WaitGroup, db *sqlx.DB, pick *models.Pick) {
	go func() {
		defer wg.Done()
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered from readability panic", r)
			}
		}()

		url := pick.Link
		timeout := time.Duration(5 * time.Second)
		client := &http.Client{
			Timeout: timeout,
		}
		res, err := client.Get(url)
		if err != nil {
			fmt.Printf("Request failed %s: %s\n", pick.Name, err)
			return
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Printf("Failed reading response body %s: %s\n", pick.Name, err)
			return
		}

		doc, err := readability.Parse(body)
		if err != nil {
			fmt.Printf("Failed parsing %s: %s\n", pick.Name, err)
			return
		}
		content, err := doc.Content()
		if err != nil {
			fmt.Printf("Failed retrieving content for %s: %s\n", pick.Name, err)
			return
		}
		trim(&content)
		pick.Content = content
		pick.UpdateContent(db)
	}()
}

func trim(content *string) {
	re := regexp.MustCompile("\\s+")
	re.ReplaceAllString(*content, " ")
	*content = strings.Replace(*content, "\n", "", -1)
}
