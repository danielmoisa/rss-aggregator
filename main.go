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

	apiConfig := api.ApiConfig{
		DB: database.New(conn),
	}

	router := chi.NewRouter()

	// Cors
	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// Init server
	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	// Api routes
	v1Router := chi.NewRouter()
	v1Router.Get("/health", api.HandlerReadiness)
	v1Router.Post("/users", apiConfig.CreateUser)
	v1Router.Get("/users", apiConfig.AuthMiddleware(apiConfig.GetUserByApiKey))

	router.Mount("/v1", v1Router)

	// Start server
	fmt.Printf("Server listen on http://localhost:%v\n", port)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
