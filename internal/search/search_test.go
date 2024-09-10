package search

import (
	"search_engine/internal/models"
	"testing"
)

func areResultsEqual(a, b []models.SearchResult) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func createIndex(articles []models.Article) *Index {
	index := NewIndex()
	for _, article := range articles {
		err := index.AddArticle(article)
		if err != nil {
			panic(err)
		}
	}
	return index
}

func TestSearch(t *testing.T) {
	tests := []struct {
		name            string
		query           string
		index           Index
		expectedResults []models.SearchResult
	}{
		{
			name:  "Single word",
			query: "sport",
			index: *createIndex([]models.Article{{ID: 0, Title: "The Benefits of sport", Content: "Regular exercise is important for maintaining good health :).", URL: "http://localhost:8080/add_article/article1"}, {ID: 1, Title: "Как управлять эффективно", Content: "Эффективное управление временем повышает продуктивность.", URL: "http://localhost:8080/add_article/article2"}}),
			expectedResults: []models.SearchResult{
				{
					Title: "The Benefits of sport",
					URL:   "http://localhost:8080/add_article/article1",
				},
			},
		},
		{
			name:  "Multiple words",
			query: "benefits of html",
			index: *createIndex([]models.Article{{ID: 0, Title: "The Benefits of sport", Content: "Regular exercise is important for maintaining good health :).", URL: "http://localhost:8080/add_article/article1"}, {ID: 1, Title: "How to Learn HTML", Content: "HTML stands for HyperText Markup Language and it is necessary to learn it.", URL: "http://localhost:8080/add_article/article4"}}),
			expectedResults: []models.SearchResult{
				{
					Title: "The Benefits of sport",
					URL:   "http://localhost:8080/add_article/article1",
				},
				{
					Title: "How to Learn HTML",
					URL:   "http://localhost:8080/add_article/article4",
				},
			},
		},
		{
			name:            "Words not present in the index",
			query:           "Benefits of Html",
			index:           *createIndex([]models.Article{{ID: 2, Title: "Как управлять эффективно", Content: "Эффективное управление временем повышает продуктивность.", URL: "http://localhost:8080/add_article/article3"}}),
			expectedResults: []models.SearchResult{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			results := tt.index.Search(tt.query)
			if len(results) != len(tt.expectedResults) {
				t.Errorf("Expected %d elements, got %d", len(tt.expectedResults), len(results))
			}
			if !areResultsEqual(results, tt.expectedResults) {
				t.Errorf("Expected %v, got %v", tt.expectedResults, results)
			}
		})
	}
}
