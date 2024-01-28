package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/danielmoisa/rss-aggregator/internal/api"
	"github.com/danielmoisa/rss-aggregator/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port not configured in the env")
	}

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("Database url not configured in the env")
	}

	conn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("Can't connect to database")
	}

	apiCfg := api.ApiConfig{
		DB: database.New(conn),
	}

	r := chi.NewRouter()

	// Cors
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// Init server
	server := &http.Server{
		Handler: r,
		Addr:    ":" + port,
	}

	// Api routes
	v1R := chi.NewRouter()
	v1R.Get("/health", api.HandlerReadiness)
	v1R.Post("/users", apiCfg.CreateUserHandler)
	v1R.Get("/users", apiCfg.AuthMiddleware(apiCfg.GetUserHandler))
	v1R.Post("/feeds", apiCfg.AuthMiddleware(apiCfg.CreateFeedHandler))
	v1R.Get("/feeds", apiCfg.GetFeedsHandler)

	r.Mount("/v1", v1R)

	// Start server
	fmt.Printf("Server listen on http://localhost:%v\n", port)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
