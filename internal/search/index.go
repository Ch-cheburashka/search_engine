package search

import (
	"search_engine/internal/models"
	"strings"
	"unicode"
)

func tokenize(content string) map[string]int {
	content = strings.ToLower(content)

	words := make(map[string]int)

	tokens := strings.FieldsFunc(content, func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	})

	for _, word := range tokens {
		if word != "" {
			words[word]++
		}
	}

	return words
}

type ArticlePair struct {
	Frequency int
	Article   models.Article
}

type PodcastPair struct {
	Frequency int
	Podcast   models.Podcast
}

type Index struct {
	Articles map[string][]ArticlePair
	Podcasts map[string][]PodcastPair
}

func NewIndex() *Index {
	return &Index{Articles: make(map[string][]ArticlePair), Podcasts: make(map[string][]PodcastPair)}
}

func (index *Index) AddArticle(article models.Article) {
	words := tokenize(article.Content + article.Title)
	for word := range words {
		if !stopWords[word] {
			index.Articles[word] = append(index.Articles[word], ArticlePair{Frequency: words[word], Article: article})
		}
	}
}

func (index *Index) AddPodcast(podcast models.Podcast) {
	words := tokenize(podcast.Description + podcast.Title)
	for word := range words {
		if !stopWords[word] {
			index.Podcasts[word] = append(index.Podcasts[word], PodcastPair{Frequency: words[word], Podcast: podcast})
		}
	}
}
