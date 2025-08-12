package main

import (
	"encoding/json"
	"log"
	"net/http"
)


func respondWithJson(w http.ResponseWriter, statusCode int, payload any) {
	data,err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON: %v", payload)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(data)
}

func respondWithError(w http.ResponseWriter, statusCode int, message string) {
	if statusCode > 499 {
		log.Printf("Server error: %v", message)
	}

	type ErrorResponse struct {
		Error string `json:"error"`
	}
	respondWithJson(w, statusCode, ErrorResponse{Error: message})
}