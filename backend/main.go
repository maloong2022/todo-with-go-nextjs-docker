package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	database "todo/internal/database"

	_ "github.com/lib/pq"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type ApiConf struct {
	DB *database.Queries
}

func main() {
	// get from docker file so no need this code now
	//err := godotenv.Load()
	//if err != nil {
	//	log.Fatal("Failed to read godotenv")
	//}

	connStr := os.Getenv("POSTGRES_URL")

	log.Println(connStr)

	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	db := database.New(conn)

	apiConf := ApiConf{
		DB: db,
	}

	r := chi.NewRouter()
	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/todos", apiConf.handlePostTodos)
	r.Delete("/todos/{id}", apiConf.handleDeleteTodos)
	r.Patch("/todos/{id}", apiConf.handleEidtTodos)
	r.Get("/todos", apiConf.handleGetAllTodos)
	r.Get(("/todos/{id}"), apiConf.handleGetTodos)

	http.ListenAndServe(":8000", r)
}
