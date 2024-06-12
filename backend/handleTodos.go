package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
	"todo/internal/database"

	"github.com/go-chi/chi/v5"
)

func (apiConf *ApiConf) handlePostTodos(w http.ResponseWriter, r *http.Request) {
	type parameter struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	params := parameter{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		log.Println("Failed to decode it %v", err)
		http.Error(w, "Error to decode it", http.StatusBadRequest)
		return
	}

	todo, err := apiConf.DB.CreateTodos(r.Context(), database.CreateTodosParams{
		Title:     params.Title,
		Content:   params.Content,
		Createdat: time.Now(),
	})
	if err != nil {
		log.Println("Failed to create todos %v", err)
		http.Error(w, "Error to create todos", http.StatusBadRequest)
		return
	}

	responseWithJSON(w, http.StatusOK, todo)
}

func (apiConf *ApiConf) handleDeleteTodos(w http.ResponseWriter, r *http.Request) {
	todoIdString := chi.URLParam(r, "id")

	todoIntId, err := strconv.ParseInt(todoIdString, 10, 64)
	if err != nil {
		log.Println("Failed to delete todos %v", err)
		http.Error(w, "Error to parse Id", http.StatusBadRequest)
		return
	}
	err = apiConf.DB.DeleteTodos(r.Context(), todoIntId)
	if err != nil {
		log.Println("Failed to delete todos %v", err)
		http.Error(w, "Error to delete todos", http.StatusBadRequest)
		return
	}

	responseWithJSON(w, http.StatusOK, map[string]string{"message": "deleted well!"})
}

func (apiConf *ApiConf) handleEidtTodos(w http.ResponseWriter, r *http.Request) {
	todoIdString := chi.URLParam(r, "id")

	todoIntId, err := strconv.ParseInt(todoIdString, 10, 64)
	if err != nil {
		log.Println("Failed to patch todos %v", err)
		http.Error(w, "Error to parse Id", http.StatusBadRequest)
		return
	}

	type parameter struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	params := parameter{}

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&params)
	if err != nil {
		log.Println("Failed to decode it %v", err)
		http.Error(w, "Error to decode it", http.StatusBadRequest)
		return
	}

	err = apiConf.DB.UpdateTodo(r.Context(), database.UpdateTodoParams{
		ID:      todoIntId,
		Title:   params.Title,
		Content: params.Content,
	})
	if err != nil {
		log.Println("Failed to update todos %v", err)
		http.Error(w, "Error to update todos", http.StatusBadRequest)
		return
	}

	responseWithJSON(w, http.StatusOK, map[string]string{"message": "successfully Edit todo!"})
}

func (apiConf *ApiConf) handleGetAllTodos(w http.ResponseWriter, r *http.Request) {
	allTodos, err := apiConf.DB.ListAllTodos(r.Context())
	if err != nil {
		log.Println("Failed to get all todos %v", err)
		http.Error(w, "Error to get all todos", http.StatusBadRequest)
		return
	}
	responseWithJSON(w, http.StatusOK, allTodos)
}

func (apiConf *ApiConf) handleGetTodos(w http.ResponseWriter, r *http.Request) {
	todoIdString := chi.URLParam(r, "id")

	todoIntId, err := strconv.ParseInt(todoIdString, 10, 64)
	if err != nil {
		log.Println("Failed to get todos %v", err)
		http.Error(w, "Error to parse Id", http.StatusBadRequest)
		return
	}

	todos, err := apiConf.DB.GetTodos(r.Context(), todoIntId)
	if err != nil {
		log.Println("Failed to get todos %v", err)
		http.Error(w, "Error to get todos", http.StatusBadRequest)
		return
	}
	responseWithJSON(w, http.StatusOK, todos)
}
