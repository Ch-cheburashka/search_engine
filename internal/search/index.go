package search

import (
	"search_engine/internal/models"
	"strings"
)

func tokenize(content string) []string {
	words := strings.Fields(strings.ToLower(content))
	return words
}

type Index struct {
	Articles map[string][]models.Article
	Podcasts map[string][]models.Podcast
}

func NewIndex() *Index {
	return &Index{Articles: make(map[string][]models.Article), Podcasts: make(map[string][]models.Podcast)}
}

func (index *Index) AddArticle(article models.Article) {
	words := tokenize(article.Content + article.Title)
	for _, word := range words {
		index.Articles[word] = append(index.Articles[word], article)
	}
}

func (index *Index) AddPodcast(podcast models.Podcast) {
	words := tokenize(podcast.Description + podcast.Title)
	for _, word := range words {
		index.Podcasts[word] = append(index.Podcasts[word], podcast)
	}
}
