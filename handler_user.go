

func (apiCfg *apiConfig) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusCreated, map[string]string{"status": "user created"})
}