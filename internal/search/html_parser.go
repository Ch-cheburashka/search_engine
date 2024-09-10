package search

import (
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
)

func ParseHTML(reader io.Reader) (Title string, Content string, err error) {
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		log.Printf("Failed to read the document : %v", err)
		return "", "", err
	}
	title := doc.Find("h1.inner-name").Text()
	content := doc.Find("div.inner-content").Text()
	return title, content, nil
}
