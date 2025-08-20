package main

import (
	"encoding/json"
	"net/http"

	"github.com/crucialjun/rss_aggregator/internal/database"
)

func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameter struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}

	var params parameter
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		Name:   params.Name,
		UserID: user.ID,
		Url:    params.Url,
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusCreated, databaseFeedToFeed(feed))
}

func (apiCfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	feedFromdb := make([]Feed, len(feeds))

	for i, feed := range feeds {
		feedFromdb[i] = databaseFeedToFeed(feed)
	}

	respondWithJson(w, http.StatusOK, feedFromdb)
}
