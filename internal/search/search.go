package search

import (
	"regexp"
	"search_engine/internal/models"
	"sort"
	"strings"
)

var stopWords = map[string]bool{
	"and": true, "or": true, "but": true, "nor": true, "so": true, "yet": true, "for": true,
	"in": true, "on": true, "at": true, "by": true, "with": true, "about": true, "against": true,
	"between": true, "into": true, "through": true, "during": true, "before": true, "after": true,
	"above": true, "below": true, "to": true, "from": true, "up": true, "down": true, "under": true,
	"over": true, "the": true, "a": true, "an": true, "he": true, "she": true, "it": true, "they": true,
	"we": true, "you": true, "I": true, "me": true, "him": true, "her": true, "us": true, "them": true,
	"is": true, "are": true, "was": true, "were": true, "be": true, "been": true, "am": true, "have": true,
	"has": true, "had": true, "do": true, "does": true, "did": true, "will": true, "would": true, "shall": true,
	"should": true, "can": true, "could": true, "may": true, "might": true, "must": true, "this": true,
	"that": true, "these": true, "those": true, "my": true, "your": true, "his": true,
	"its": true, "our": true, "their": true, "of": true, "if": true, "then": true, "there": true,
	"here": true, "when": true, "where": true, "why": true, "how": true, "which": true, "no": true,
	"not": true, "neither": true, "never": true, "none": true, "very": true, "too": true, "quite": true,
	"rather": true, "almost": true, "just": true, "only": true,
}

func (index *Index) Search(query string) []models.SearchResult {
	query = strings.ToLower(query)
	re := regexp.MustCompile(`[^\w\s]+`)
	cleanedContent := re.ReplaceAllString(query, "")
	words := strings.Fields(cleanedContent)

	searchResults := make(map[string]models.SearchResult)
	for _, word := range words {
		if stopWords[word] {
			continue
		}
		articles := index.Articles[word]
		podcasts := index.Podcasts[word]

		sort.Slice(articles, func(i, j int) bool {
			return articles[i].Article.ID < articles[j].Article.ID
		})

		sort.Slice(articles, func(i, j int) bool {
			return articles[i].Article.ID < articles[j].Article.ID
		})

		for _, article := range articles {
			searchResults[article.Article.URL] = models.SearchResult{Title: article.Article.Title, URL: article.Article.URL}
		}
		for _, podcast := range podcasts {
			searchResults[podcast.Podcast.URL] = models.SearchResult{Title: podcast.Podcast.Title, URL: podcast.Podcast.URL}
		}
	}

	results := make([]models.SearchResult, 0)

	for _, result := range searchResults {
		results = append(results, result)
	}

	return results
}
