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

func (apiConfig *ApiConfig) CreateFeedHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		utils.ResponseWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	feed, err := apiConfig.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	})
	if err != nil {
		utils.ResponseWithError(w, 400, fmt.Sprintf("Couldn't create a new feed: %v", err))
	}

	utils.ResponseWithJSON(w, 200, feed)
}

func (apiConfig *ApiConfig) GetFeedsHandler(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiConfig.DB.GetFeeds(r.Context())
	if err != nil {
		utils.ResponseWithError(w, 400, fmt.Sprintf("Couldn't create feed: %v", err))
		return
	}

	utils.ResponseWithJSON(w, 201, feeds)
}
