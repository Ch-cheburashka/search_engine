package search

import "search_engine/internal/models"

func (index *Index) Search(query string) []models.SearchResult {
	words := tokenize(query)
	searchResults := make([]models.SearchResult, 0)
	for _, word := range words {
		articles := index.Articles[word]
		podcasts := index.Podcasts[word]

		for i := range articles {
			searchResults = append(searchResults, models.SearchResult{Title: articles[i].Title, URL: articles[i].URL})
		}
		for i := range podcasts {
			searchResults = append(searchResults, models.SearchResult{Title: podcasts[i].Title, URL: podcasts[i].URL})
		}
	}
	return searchResults
}
