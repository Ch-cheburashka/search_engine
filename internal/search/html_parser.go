package search

import (
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
)

func ParseHTML(URL string) (Title string, Content string, err error) {
	res, err := http.Get(URL)
	if err != nil {
		log.Printf("Failed to fetch URL %s: %v", URL, err)
		return "", "", err
	}
	defer func(Body io.ReadCloser) {
		if err := res.Body.Close(); err != nil {
			log.Printf("Failed to close response body: %v", err)
		}
	}(res.Body)

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Printf("Failed to parse HTML from URL %s: %v", URL, err)
		return "", "", err
	}
	title := doc.Find("h1.title").Text()
	content := doc.Find("div.content").Text()

	return title, content, nil
}
