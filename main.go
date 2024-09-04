package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"search_engine/internal/models"
	"search_engine/internal/search"
	"strings"
	"sync"
)

var index = search.NewIndex()
var articlesNumber = 0
var podcastsNumber = 0
var artMux sync.Mutex
var podMux sync.Mutex

func addArticleHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		log.Printf("Method \"%s\" Not Allowed\n", request.Method)
		http.Error(writer, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	title, content, err := search.ParseHTML(request.Body)
	if err != nil {
		log.Println("Failed to parse the document")
		http.Error(writer, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if len(title) == 0 || len(content) == 0 {
		log.Println("Missing title or content")
		http.Error(writer, "Missing title or content", http.StatusBadRequest)
		return
	}
	artMux.Lock()
	article := models.Article{ID: articlesNumber + 1, Title: title, Content: content, URL: request.URL.Path + "/" + strings.ReplaceAll(title, " ", "_")}
	articlesNumber++
	artMux.Unlock()

	index.AddArticle(article)
	writer.WriteHeader(http.StatusOK)
	_, err = writer.Write([]byte("Article added successfully"))
	if err != nil {
		log.Println("Failed to write response")
		http.Error(writer, "Internal server error", http.StatusInternalServerError)
		return
	}

}

func addPodcastHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		log.Printf("Method \"%s\" Not Allowed\n", request.Method)
		http.Error(writer, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	title, description, err := search.ParseHTML(request.Body)
	if err != nil {
		log.Println("Failed to parse the document")
		http.Error(writer, "Invalid request payload", http.StatusBadRequest)
		return
	}
	podMux.Lock()
	podcast := models.Podcast{ID: podcastsNumber + 1, Title: title, Description: description, URL: request.URL.Path + "/" + strings.ReplaceAll(title, " ", "_")}
	if len(title) == 0 || len(description) == 0 {
		log.Println("Missing title or description")
		http.Error(writer, "Missing title or description", http.StatusBadRequest)
		return
	}
	podcastsNumber++
	podMux.Unlock()

	index.AddPodcast(podcast)
	writer.WriteHeader(http.StatusOK)
	_, err = writer.Write([]byte("Podcast added successfully"))
	if err != nil {
		log.Println("Failed to write response")
		http.Error(writer, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func searchHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		log.Printf("Method \"%s\" Not Allowed\n", request.Method)
		http.Error(writer, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	query := request.URL.Query().Get("query")
	if query == "" {
		log.Println("Query parameter is missing")
		http.Error(writer, "Invalid query", http.StatusBadRequest)
		return
	}
	results := index.Search(query)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	jsonData, err := json.Marshal(results)
	if err != nil {
		log.Println("Failed to write create JSON object")
		http.Error(writer, "Internal server error", http.StatusInternalServerError)
		return
	}
	_, err = writer.Write(jsonData)
	if err != nil {
		log.Println("Failed to write response")
		http.Error(writer, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func main() {
	port := flag.String("port", "8080", "Port to listen on")
	flag.Parse()

	http.HandleFunc("/search", searchHandler)
	http.HandleFunc("/add_article", addArticleHandler)
	http.HandleFunc("/add_podcast", addPodcastHandler)

	log.Printf("Server started at http://localhost:%s\n", *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
