package main

import (
	"encoding/json"
	"net/http"

	"github.com/crucialjun/rss_aggregator/internal/database"
)

func (apiCfg *apiConfig) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameter struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	var params parameter
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		Name:  params.Name,
		Email: params.Email,
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	data := databaseUserToUser(user)

	respondWithJson(w, http.StatusCreated, data)
}

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {

	respondWithJson(w, http.StatusOK, databaseUserToUser(user))
}
