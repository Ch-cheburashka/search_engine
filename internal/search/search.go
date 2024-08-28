package search

import (
	"regexp"
	"search_engine/internal/models"
	"strings"
)

func (index *Index) Search(query string) []models.SearchResult {
	strings.ToLower(query)
	re := regexp.MustCompile(`[^\w\s]+`)
	cleanedContent := re.ReplaceAllString(query, "")
	words := strings.Fields(cleanedContent)
	searchResults := make(map[string]models.SearchResult)
	for _, word := range words {
		articles := index.Articles[word]
		podcasts := index.Podcasts[word]
		for _, article := range articles {
			searchResults[article.URL] = models.SearchResult{Title: article.Title, URL: article.URL}
		}
		for _, podcast := range podcasts {
			searchResults[podcast.URL] = models.SearchResult{Title: podcast.Title, URL: podcast.URL}
		}
	}

	results := make([]models.SearchResult, 0)

	for _, result := range searchResults {
		results = append(results, result)
	}

	return results
}
