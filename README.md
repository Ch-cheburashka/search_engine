# Documentation

A simple search engine that allows you to add articles and podcasts, and search for them. Built using Go.

## Table of Contents
1. [Usage](#usage)
2. [API endpoints](#api-endpoints)
3. [Data Models](#data-models)

## Usage

### Prerequisites

- **Go**: Ensure Go (version 1.16 or later) is installed on your system.
### Steps

1. **Clone the Repository**
    ```bash
    git clone https://github.com/Ch-cheburashka/search_engine
    cd search_engine
    ```
2. **Install Dependencies**
   Run the following command to install any required Go modules:
    ```bash
    go mod tidy
    ```

3. **Run the Server**
   To start the server, use the following command: 
    ```bash
    go run main.go
   ```

The server will start on http://localhost:8080 ***by default***. You can specify a different port using the `-port=... flag`.

## API Endpoints 
***1. Add an Article***
  -  Endpoint: `/add_article`
  - Method: `POST`
  - Request: HTML content where the title is read from `<h1 class="inner-title">` and the content from `<div class="inner-content">`.
  - Response:
    - 200 OK: Article added successfully.

***2. Add a Podcast***
 - Endpoint: `/add_podcast`
 - Method: `POST`
 - Request: HTML content where the title is read from `<h1 class="inner-title">` and the content from `<div class="inner-content">`.
 - Response:
   - 200 OK: Article added successfully.

***3. Search***
 - Endpoint: `/search`
 - Method: `GET`
 - Query Parameter: query - the search query string. `http://localhost:8080/search?query=word`
 - Response:
     - 200 OK: Returns a JSON array of search results. Each result is an object with the following format:
       `[ { "title": "Example Title", "url": "http://example.com" }, ... ]`


## Data Models
### Article
- ID: Unique identifier for the article (the number of added article).
- Title
- Content
- URL (need to be refactored)

### Podcast
- ID: Unique identifier for the podcast (the number of added podcast).
- Title
- Description
- URL (need to be refactored)

### Search Result
- Title
- URL (need to be refactored)