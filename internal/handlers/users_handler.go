package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/danielmoisa/rss-aggregator/internal/database"
	"github.com/danielmoisa/rss-aggregator/pkg/response"
	"github.com/google/uuid"
)

type ApiConfig struct {
	DB *database.Queries
}

func (apiConfig *ApiConfig) CreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json: "name"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		response.ResponseWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	user, err := apiConfig.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		response.ResponseWithError(w, 400, fmt.Sprintf("Couldn't create a new user: %v", err))
	}

	response.ResponseWithJSON(w, 200, user)
}
