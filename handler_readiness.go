package main

import (
	"log"
	"net/http"
)

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	log.Println("Readiness probe requested")
	respondWithJson(w, http.StatusOK, map[string]string{"status": "ready"})
}
