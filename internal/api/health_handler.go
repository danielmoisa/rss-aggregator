package api

import (
	"net/http"

	"github.com/danielmoisa/rss-aggregator/internal/utils"
)

func HandlerReadiness(w http.ResponseWriter, r *http.Request) {
	utils.ResponseWithJSON(w, 200, struct{}{})
}
