package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/danielmoisa/rss-aggregator/internal/database"
	"github.com/danielmoisa/rss-aggregator/internal/utils"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (apiCfg *ApiConfig) CreateFeedFollowsHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedId uuid.UUID `json:"feedId"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		utils.ResponseWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	feedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedId,
	})
	if err != nil {
		utils.ResponseWithError(w, 400, fmt.Sprintf("Couldn't create a new feed follow: %v", err))
	}

	utils.ResponseWithJSON(w, 200, feedFollow)
}

func (apiCfg *ApiConfig) GetFeedsFollowsHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	feed, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		utils.ResponseWithError(w, 400, fmt.Sprintf("Couldn't get feed follows: %v", err))
		return
	}

	utils.ResponseWithJSON(w, 201, feed)
}

func (apiCfg *ApiConfig) DeleteFeedsFollowHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIdStr := chi.URLParam(r, "feedFollowId")
	feedFollowId, err := uuid.Parse(feedFollowIdStr)
	if err != nil {
		utils.ResponseWithError(w, 400, fmt.Sprintf("Couldn't parse feed follows: %v", err))
		return
	}

	err = apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     feedFollowId,
		UserID: user.ID,
	})
	if err != nil {
		utils.ResponseWithError(w, 400, fmt.Sprintf("Couldn't delete feed follows: %v", err))
		return
	}

	utils.ResponseWithJSON(w, 200, struct{}{})
}
