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
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/todos", apiConf.handlePostTodos)
	r.Delete("/todos/{id}", apiConf.handleDeleteTodos)
	r.Patch("/todos/{id}", apiConf.handleEidtTodos)
	r.Get("/todos", apiConf.handleGetAllTodos)
	r.Get(("/todos/{id}"), apiConf.handleGetTodos)

	http.ListenAndServe(":8000", r)
}
