package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get ENV vars here

	// Initialize App Dependencies here

	// Initialize Server
	mux := http.NewServeMux()
	SetupRoutes(mux)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	log.Println("Listening on port 8080")
	server.ListenAndServe()
}
