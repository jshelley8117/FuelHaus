package main

import (
	"context"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/jshelley8117/FuelHaus/internal/resource"
)

func main() {
	ctx := context.Background()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get ENV vars here

	// Initialize App Dependencies here
	firebaseServices, err := resource.InitializeFirebaseServices(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer firebaseServices.Firestore.Close()

	// Initialize Server
	mux := http.NewServeMux()
	SetupRoutes(mux, *firebaseServices)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	log.Println("Listening on port 8080")
	server.ListenAndServe()
}
