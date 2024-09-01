package main

import (
	"encoding/json"
	"log"
	"net/http"
	"search_engine/internal/models"
	"search_engine/internal/search"
)

var index = search.NewIndex()
var articlesNumber = 0
var podcastsNumber = 0

func addArticleHandler(writer http.ResponseWriter, request *http.Request) {
	title, content, err := search.ParseHTML(request.Body)
	if err != nil {
		http.Error(writer, "Invalid request payload", http.StatusBadRequest)
		return
	}
	article := models.Article{ID: articlesNumber + 1, Title: title, Content: content, Author: "John Doe", URL: request.URL.Path + "/" + title}
	index.AddArticle(article)
	writer.WriteHeader(http.StatusOK)
	_, err = writer.Write([]byte("Article added successfully"))
	if err != nil {
		http.Error(writer, "Internal server error", http.StatusInternalServerError)
		return
	}
	articlesNumber++
}

func addPodcastHandler(writer http.ResponseWriter, request *http.Request) {
	title, description, err := search.ParseHTML(request.Body)
	if err != nil {
		http.Error(writer, "Invalid request payload", http.StatusBadRequest)
	}
	podcast := models.Podcast{ID: podcastsNumber + 1, Title: title, Description: description, URL: request.URL.Path + "/" + title}
	index.AddPodcast(podcast)
	writer.WriteHeader(http.StatusOK)
	_, err = writer.Write([]byte("Podcast added successfully"))
	if err != nil {
		http.Error(writer, "Internal server error", http.StatusInternalServerError)
		return
	}
	podcastsNumber++
}

func searchHandler(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query().Get("query")
	if query == "" {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	results := index.Search(query)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	jsonData, err := json.Marshal(results)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = writer.Write(jsonData)
	if err != nil {
		http.Error(writer, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/search", searchHandler)
	http.HandleFunc("/add_article", addArticleHandler)
	http.HandleFunc("/add_podcast", addPodcastHandler)

	log.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
