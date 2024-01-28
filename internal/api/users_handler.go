package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/danielmoisa/rss-aggregator/internal/database"
	"github.com/danielmoisa/rss-aggregator/internal/utils"
	"github.com/google/uuid"
)

func (apiCfg *ApiConfig) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		utils.ResponseWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		utils.ResponseWithError(w, 400, fmt.Sprintf("Couldn't create a new user: %v", err))
	}

	utils.ResponseWithJSON(w, 200, user)
}

func (apiCfg *ApiConfig) GetUserHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	utils.ResponseWithJSON(w, 200, user)
}
