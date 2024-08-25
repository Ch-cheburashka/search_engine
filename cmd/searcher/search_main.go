package main

import (
	"fmt"
	"log"
	"search_engine/internal/models"
	"search_engine/internal/search"
)

func main() {
	index := search.NewIndex()

	title1, content1, err := search.ParseHTML("http://localhost:8080/test.html")
	if err != nil {
		log.Fatal(err)
	}
	title2, content2, err := search.ParseHTML("http://localhost:8080/test_digital.html")
	if err != nil {
		log.Fatal(err)
	}
	title3, content3, err := search.ParseHTML("http://localhost:8080/podcast.html")
	if err != nil {
		log.Fatal(err)
	}
	article1 := models.Article{ID: 1, Title: title1, Content: content1, Author: "John Doe", URL: "http://localhost:8080/test.html"}
	article2 := models.Article{ID: 2, Title: title2, Content: content2, Author: "Jane Doe", URL: "http://localhost:8080/test_digital.html"}

	podcast1 := models.Podcast{ID: 1, Title: title3, Description: content3, URL: "http://localhost:8080/podcast.html"}

	index.AddPodcast(podcast1)

	index.AddArticle(article1)
	index.AddArticle(article2)

	results := index.Search("rapidly")

	for _, result := range results {
		fmt.Println(result)
	}

}
