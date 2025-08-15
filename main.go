package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/crucialjun/rss_aggregator/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // PostgreSQL driver
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	fmt.Println("Hello, World!")

	godotenv.Load()

	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal(
			"PORT environment variable is not set",
		)
	}

	dbURL := os.Getenv("DB_URL")

	if dbURL == "" {
		log.Fatal(
			"DB_URL environment variable is not set",
		)
	}

	conn, err := sql.Open("postgres", dbURL)

	if err != nil {
		log.Fatal(
			"Failed to connect to database:",
			err,
		)
	}

	queries, err := database.New(conn)
	if err != nil {
		log.Fatal(
			"Failed to create database queries:",
			err,
		)
	}



	apiCfg := apiConfig{
		DB: queries,
	
	}



	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerError)
	v1Router.Post("/users", apiCfg.handleCreateUser)

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
