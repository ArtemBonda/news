package main

import (
	"log"
	"net/http"
	"os"

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

	fs := http.FileServer(http.Dir("assets"))
	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	mux.HandleFunc("/", handlers.IndexHandler)

	log.Fatalln(http.ListenAndServe(":"+port, mux))
}
