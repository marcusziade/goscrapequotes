package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gocolly/colly"
)

type Quote struct {
	Quote  string `json:"quote"`
	Author string `json:"author"`
}

var fetcher = fetchAllQuotes

func main() {
	http.HandleFunc("/quotes", QuotesHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func QuotesHandler(w http.ResponseWriter, r *http.Request) {
	quotes, err := fetcher()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.MarshalIndent(quotes, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func fetchAllQuotes() ([]Quote, error) {
	var allQuotes []Quote
	page := 1

	for {
		quotes, err := fetchQuotes(page)
		if err != nil {
			return nil, err
		}

		if len(quotes) == 0 {
			break
		}

		allQuotes = append(allQuotes, quotes...)
		page++
	}

	return allQuotes, nil
}

func fetchQuotes(page int) ([]Quote, error) {
	url := fmt.Sprintf("http://quotes.toscrape.com/page/%d/", page)
	c := colly.NewCollector()
	var quotesList []Quote

	c.OnHTML("div.quote", func(e *colly.HTMLElement) {
		quote := e.ChildText("span.text")
		author := e.ChildText("small.author")
		quotesList = append(quotesList, Quote{Quote: quote, Author: author})
	})

	err := c.Visit(url)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch quotes from page %d: %w", page, err)
	}

	return quotesList, nil
}
