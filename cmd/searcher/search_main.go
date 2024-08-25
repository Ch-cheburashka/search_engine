package main

import (
	"fmt"
	"search_engine/internal/models"
	"search_engine/internal/search"
)

func main() {
	index := search.NewIndex()

	article1 := models.Article{ID: 1, Title: "Go Language", Content: "Discover tips, techniques, and discussions on Go programming.", Author: "John Doe", URL: "/articles/go-language"}
	article2 := models.Article{ID: 2, Title: "Digital News", Content: "Digital news is evolving", Author: "Jane Doe", URL: "/articles/digital-news"}

	podcast1 := models.Podcast{ID: 1, Title: "Mental struggles", Description: "Exploring mental health through personal stories and expert insights", URL: "/podcasts/mental-struggles"}

	index.AddArticle(article1)
	index.AddArticle(article2)

	index.AddPodcast(podcast1)

	results := index.Search("and")

	for _, result := range results {
		fmt.Println(result)
	}
}
