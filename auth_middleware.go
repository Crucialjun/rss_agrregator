package main

import (
	"net/http"

	"github.com/crucialjun/rss_aggregator/internal/auth"
	"github.com/crucialjun/rss_aggregator/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) authMiddleware(next authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apikey, error := auth.GetApiKey(r.Header)

		if error != nil {
			respondWithError(w, http.StatusUnauthorized, error.Error())
			return
		}

		user, err := apiCfg.DB.GetUser(r.Context(), apikey)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		next(w, r, user)
	}

}
