package search

import (
	"search_engine/internal/models"
	"testing"
)

func areArticlesEqual(a, b []ArticlePair) bool {
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

func arePodcastsEqual(a, b []PodcastPair) bool {
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

var englishArticle1 = models.Article{ID: 0, Title: "The Benefits of sport", Content: "Regular exercise is important for maintaining good health :).", URL: "http://localhost:8080/add_article/article1"}
var englishArticle2 = models.Article{ID: 1, Title: "The Benefits of sport", Content: "Regular exercise is important for maintaining good health :).", URL: "http://localhost:8080/add_article/article2"}
var russianArticle = models.Article{ID: 2, Title: "Как управлять эффективно", Content: "Эффективное управление временем повышает продуктивность.", URL: "http://localhost:8080/add_article/article3"}
var frequentWordsArticle = models.Article{ID: 3, Title: "How to Learn HTML", Content: "HTML stands for HyperText Markup Language and it is necessary to learn it.", URL: "http://localhost:8080/add_article/article4"}

func TestIndex_AddArticle(t *testing.T) {
	tests := []struct {
		name        string
		articles    []models.Article
		expectedMap map[string][]ArticlePair
		expectError bool
	}{
		{
			name: "Unique set of words in articles",
			articles: []models.Article{
				englishArticle1,
				russianArticle,
			},
			expectedMap: map[string][]ArticlePair{
				"benefits": {
					{1, englishArticle1},
				},
				"regular": {
					{1, englishArticle1},
				},
				"exercise": {
					{1, englishArticle1},
				},
				"sport": {
					{1, englishArticle1},
				},
				"important": {
					{1, englishArticle1},
				},
				"maintaining": {
					{1, englishArticle1},
				},
				"good": {
					{1, englishArticle1},
				},
				"health": {
					{1, englishArticle1},
				},
				"как": {
					{1, russianArticle},
				},
				"управлять": {
					{1, russianArticle},
				},
				"эффективно": {
					{1, russianArticle},
				},
				"эффективное": {
					{1, russianArticle},
				},
				"управление": {
					{1, russianArticle},
				},
				"временем": {
					{1, russianArticle},
				},
				"повышает": {
					{1, russianArticle},
				},
				"продуктивность": {
					{1, russianArticle},
				},
			},
			expectError: false,
		},
		{
			name: "Overlapping words in articles",
			articles: []models.Article{
				englishArticle1,
				englishArticle2,
			},
			expectedMap: map[string][]ArticlePair{
				"benefits": {
					{1, englishArticle1},
					{1, englishArticle2},
				},
				"regular": {
					{1, englishArticle1},
					{1, englishArticle2},
				},
				"exercise": {
					{1, englishArticle1},
					{1, englishArticle2},
				},
				"sport": {
					{1, englishArticle1},
					{1, englishArticle2},
				},
				"important": {
					{1, englishArticle1},
					{1, englishArticle2},
				},
				"maintaining": {
					{1, englishArticle1},
					{1, englishArticle2},
				},
				"good": {
					{1, englishArticle1},
					{1, englishArticle2},
				},
				"health": {
					{1, englishArticle1},
					{1, englishArticle2},
				},
			},
			expectError: false,
		},
		{
			name:     "Frequent words in articles",
			articles: []models.Article{frequentWordsArticle},
			expectedMap: map[string][]ArticlePair{
				"learn": {
					{2, frequentWordsArticle},
				},
				"html": {
					{2, frequentWordsArticle},
				},
				"stands": {
					{1, frequentWordsArticle},
				},
				"hypertext": {
					{1, frequentWordsArticle},
				},
				"markup": {
					{1, frequentWordsArticle},
				},
				"language": {
					{1, frequentWordsArticle},
				},
				"necessary": {
					{1, frequentWordsArticle},
				},
			},
			expectError: false,
		},
		{
			name: "Empty elements in articles",
			articles: []models.Article{
				{ID: 0, Title: "", Content: "", URL: ""},
			},
			expectedMap: map[string][]ArticlePair{},
			expectError: true,
		},
	}
	for _, tt := range tests {
		index := NewIndex()
		t.Run(tt.name, func(t *testing.T) {
			for _, article := range tt.articles {
				err := index.AddArticle(article)
				if tt.expectError && err == nil {
					t.Errorf("Expected error but got none")
				} else if !tt.expectError && err != nil {
					t.Errorf("Did not expect error but got one: %v", err)
				}
			}
			if len(index.Articles) != len(tt.expectedMap) {
				t.Errorf("Expected %d elements, got %d", len(tt.expectedMap), len(index.Articles))
			}
			for word := range tt.expectedMap {
				if _, found := index.Articles[word]; !found {
					t.Errorf("Expected word %q not found in index", word)
				}
			}
			for word := range index.Articles {
				if !areArticlesEqual(index.Articles[word], tt.expectedMap[word]) {
					t.Errorf("Expected %v frequency, got %v", tt.expectedMap[word], index.Articles[word])
				}
			}

		})
	}
}

var englishPodcast1 = models.Podcast{ID: 0, Title: "Introduction to Frontend", Description: "CSS stands for Cascading Style Sheets and it is used to layout webpages.", URL: "http://localhost:8080/add_podcast/podcast1"}
var englishPodcast2 = models.Podcast{ID: 1, Title: "Introduction to Frontend", Description: "CSS stands for Cascading Style Sheets and it is used to layout webpages.", URL: "http://localhost:8080/add_podcast/podcast2"}
var russianPodcast = models.Podcast{ID: 2, Title: "Преимущества ЗОЖ", Description: "Здоровое питание поддерживает хорошее самочувствие.", URL: "http://localhost:8080/add_podcast/podcast3"}
var frequentWordsPodcast = models.Podcast{ID: 3, Title: "How to Learn HTML", Description: "HTML stands for HyperText Markup Language and it is necessary to learn it.", URL: "http://localhost:8080/add_podcast/podcast4"}

func TestIndex_AddPodcast(t *testing.T) {
	tests := []struct {
		name        string
		podcasts    []models.Podcast
		expectedMap map[string][]PodcastPair
		expectError bool
	}{
		{
			name: "Unique set of words in podcasts",
			podcasts: []models.Podcast{
				englishPodcast1,
				russianPodcast,
			},
			expectedMap: map[string][]PodcastPair{
				"introduction": {
					{1, englishPodcast1},
				},
				"frontend": {
					{1, englishPodcast1},
				},
				"css": {
					{1, englishPodcast1},
				},
				"stands": {
					{1, englishPodcast1},
				},
				"cascading": {
					{1, englishPodcast1},
				},
				"style": {
					{1, englishPodcast1},
				},
				"sheets": {
					{1, englishPodcast1},
				},
				"used": {
					{1, englishPodcast1},
				},
				"layout": {
					{1, englishPodcast1},
				},
				"webpages": {
					{1, englishPodcast1},
				},
				"преимущества": {
					{1, russianPodcast},
				},
				"зож": {
					{1, russianPodcast},
				},
				"здоровое": {
					{1, russianPodcast},
				},
				"питание": {
					{1, russianPodcast},
				},
				"поддерживает": {
					{1, russianPodcast},
				},
				"хорошее": {
					{1, russianPodcast},
				},
				"самочувствие": {
					{1, russianPodcast},
				},
			},
			expectError: false,
		},
		{
			name: "Overlapping words in podcasts",
			podcasts: []models.Podcast{
				englishPodcast1,
				englishPodcast2,
			},
			expectedMap: map[string][]PodcastPair{
				"introduction": {
					{1, englishPodcast1},
					{1, englishPodcast2},
				},
				"frontend": {
					{1, englishPodcast1},
					{1, englishPodcast2},
				},
				"css": {
					{1, englishPodcast1},
					{1, englishPodcast2},
				},
				"stands": {
					{1, englishPodcast1},
					{1, englishPodcast2},
				},
				"cascading": {
					{1, englishPodcast1},
					{1, englishPodcast2},
				},
				"style": {
					{1, englishPodcast1},
					{1, englishPodcast2},
				},
				"sheets": {
					{1, englishPodcast1},
					{1, englishPodcast2},
				},
				"used": {
					{1, englishPodcast1},
					{1, englishPodcast2},
				},
				"layout": {
					{1, englishPodcast1},
					{1, englishPodcast2},
				},
				"webpages": {
					{1, englishPodcast1},
					{1, englishPodcast2},
				},
			},
			expectError: false,
		},
		{
			name:     "Frequent words in podcasts",
			podcasts: []models.Podcast{frequentWordsPodcast},
			expectedMap: map[string][]PodcastPair{
				"learn": {
					{2, frequentWordsPodcast},
				},
				"html": {
					{2, frequentWordsPodcast},
				},
				"stands": {
					{1, frequentWordsPodcast},
				},
				"hypertext": {
					{1, frequentWordsPodcast},
				},
				"markup": {
					{1, frequentWordsPodcast},
				},
				"language": {
					{1, frequentWordsPodcast},
				},
				"necessary": {
					{1, frequentWordsPodcast},
				},
			},
			expectError: false,
		},
		{
			name: "Empty elements in podcasts",
			podcasts: []models.Podcast{
				{ID: 0, Title: "", Description: "", URL: ""},
			},
			expectedMap: map[string][]PodcastPair{},
			expectError: true,
		},
	}
	for _, tt := range tests {
		index := NewIndex()
		t.Run(tt.name, func(t *testing.T) {
			for _, podcast := range tt.podcasts {
				err := index.AddPodcast(podcast)
				if tt.expectError && err == nil {
					t.Errorf("Expected error but got none")
				} else if !tt.expectError && err != nil {
					t.Errorf("Did not expect error but got one: %v", err)
				}
			}
			if len(index.Podcasts) != len(tt.expectedMap) {
				t.Errorf("Expected %d elements, got %d", len(tt.expectedMap), len(index.Podcasts))
			}
			for word := range tt.expectedMap {
				if _, found := index.Podcasts[word]; !found {
					t.Errorf("Expected word %q not found in index", word)
				}
			}
			for word := range index.Articles {
				if !arePodcastsEqual(index.Podcasts[word], tt.expectedMap[word]) {
					t.Errorf("Expected %v frequency, got %v", tt.expectedMap[word], index.Podcasts[word])
				}
			}

		})
	}
}
