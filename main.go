package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello, World!")

	godotenv.Load()

	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal(
			"PORT environment variable is not set",
		)
	}

	router := chi.NewRouter()

	server := &http.Server{
		Addr:    ":" + portString,
		Handler: router,
	}

	err := server.ListenAndServe()

	log.Println("Starting server on port:", portString)

	if err != nil {
		log.Fatal(err)
	}

}
