package api

import "github.com/danielmoisa/rss-aggregator/internal/database"

type ApiConfig struct {
	DB *database.Queries
}
