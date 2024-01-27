package handlers

import (
	"net/http"

	"github.com/danielmoisa/rss-aggregator/pkg/response"
)

func HandlerReadiness(w http.ResponseWriter, r *http.Request) {
	response.ResponseWithJSON(w, 200, struct{}{})
}
