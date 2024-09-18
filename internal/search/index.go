package search

import (
	"errors"
	"github.com/Ch-cheburashka/search_engine/internal/models"
	"strings"
	"sync"
	"unicode"
)

func tokenize(content string) map[string]int {
	content = strings.ToLower(content)

	words := make(map[string]int)

	tokens := strings.FieldsFunc(content, func(c rune) bool {
		return !unicode.IsLetter(c)
	})

	for _, word := range tokens {
		if word != "" && !stopWords[word] {
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
	Articles     map[string][]ArticlePair
	Podcasts     map[string][]PodcastPair
	articleByURL map[string]models.Article
	podcastByURL map[string]models.Podcast
	artMux       *sync.Mutex
	podMux       *sync.Mutex
}

func NewIndex() *Index {
	return &Index{Articles: make(map[string][]ArticlePair), Podcasts: make(map[string][]PodcastPair), articleByURL: make(map[string]models.Article), podcastByURL: make(map[string]models.Podcast), artMux: &sync.Mutex{}, podMux: &sync.Mutex{}}
}

func (index *Index) AddArticle(article models.Article) error {
	if article.Title == "" || article.Content == "" || article.URL == "" {
		return errors.New("article is empty")
	}
	words := tokenize(article.Content + article.Title)
	index.artMux.Lock()
	if _, ok := index.articleByURL[article.URL]; ok {
		index.artMux.Unlock()
		return errors.New("article already exists")
	}
	index.articleByURL[article.URL] = article
	defer index.artMux.Unlock()
	for word := range words {
		if !stopWords[word] {
			index.Articles[word] = append(index.Articles[word], ArticlePair{Frequency: words[word], Article: article})
		}
	}
	return nil
}

func (index *Index) AddPodcast(podcast models.Podcast) error {
	if podcast.Title == "" || podcast.Description == "" || podcast.URL == "" {
		return errors.New("podcast is empty")
	}
	if _, ok := index.Podcasts[podcast.URL]; ok {
		return errors.New("podcast already exists")
	}
	index.podcastByURL[podcast.URL] = podcast
	words := tokenize(podcast.Description + podcast.Title)
	index.podMux.Lock()
	defer index.podMux.Unlock()
	for word := range words {
		if !stopWords[word] {
			index.Podcasts[word] = append(index.Podcasts[word], PodcastPair{Frequency: words[word], Podcast: podcast})
		}
	}
	return nil
}
