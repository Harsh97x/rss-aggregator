package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pg"

	"github.com/harsh97x/rss-aggregator/internal/database"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load(".env")
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not found in the environment")
	}
	fmt.Println("DB_URL", dbURL)

	conn, err := sql.Open("postgresql", dbURL)
	if err != nil {
		log.Fatal("can't connect to database:", err)
	}
	apiCfg := apiConfig{
		DB: database.New(conn),
	}

	router := chi.NewRouter()

	router.Use(
		cors.Handler(
			cors.Options{
				AllowedOrigins: []string{"https://*", "http://*"},
				AllowedMethods: []string{
					"GET", "POST", "PUT", "DELETE", "OPTIONS",
				},
				AllowedHeaders:   []string{"*"},
				ExposedHeaders:   []string{"Link"},
				AllowCredentials: false,
				MaxAge:           300,
			},
		),
	)

	v1router := chi.NewRouter()
	v1router.Get("/healthz", handlerReadiness)
	v1router.Get("/err", handlerErr)
	v1router.Post("/user", apiCfg.handlerCreateUser)

	router.Mount("/v1", v1router)

	srv := &http.Server{Handler: router, Addr: ":" + portString}
	log.Printf("Server is starting at Port: %v", portString)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
