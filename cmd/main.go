package main

import (
	"errors"
	"github.com/ArtemBonda/news/internal/news"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"

	"github.com/ArtemBonda/news/internal/handlers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	apiKey := os.Getenv("NEWS_API_KEY")
	if apiKey == "" {
		log.Fatalln("Env: apiKey must be set")
	}

	newsClient := &http.Client{
		Timeout: 2 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			if len(via) >= 5 {
				return errors.New("stopped after five redirects")
			}
			return nil
		},
	}
	newsAPI := news.NewClient(newsClient, apiKey, 20)

	fs := http.FileServer(http.Dir("assets"))
	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	mux.HandleFunc("/search", handlers.SearchMiddleware(newsAPI))
	mux.HandleFunc("/", handlers.IndexHandler)

	log.Fatalln(http.ListenAndServe(":"+port, mux))
}
