package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/silvestrevivo/complete-go-webserver/internal/database"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	// we load the .env file
	godotenv.Load(".env")

	// we get the port
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT environment variable not set")
	}

	// we get the database URL
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL environment variable not set")
	}

	// Open the database connection
	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Cant open the database connection")
	}

	// Create the API router
	router := chi.NewRouter()

	apiCfg := apiConfig{
		DB: database.New(conn),
	}

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		MaxAge:           300,
		Debug:            true,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", ReadinessHandler)
	v1Router.Get("/err", ErrorHandler)
	v1Router.Post("/users", apiCfg.CreateUserHandler)

	router.Mount("/v1", v1Router)

	// Server listening
	srv := &http.Server{
		Addr:    ":" + portString,
		Handler: router,
	}

	log.Printf("Server running on PORT: %s", portString)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
