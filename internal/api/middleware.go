package api

import (
	"fmt"
	"net/http"

	"github.com/danielmoisa/rss-aggregator/internal/database"
	"github.com/danielmoisa/rss-aggregator/internal/utils"
)

type AuthHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiConfig *ApiConfig) AuthMiddleware(handler AuthHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := getApiKey(r.Header)
		if err != nil {
			utils.ResponseWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
			return
		}

		user, err := apiConfig.DB.GetUserByApiKey(r.Context(), apiKey)
		if err != nil {
			utils.ResponseWithError(w, 404, fmt.Sprintf("Couldn't get user: %v", err))
			return
		}

		handler(w, r, user)
	}
}
