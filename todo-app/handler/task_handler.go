package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/rahulp18/todo/middleware"
	"github.com/rahulp18/todo/service"
)

type ResponseType struct {
	Message string `json:"message"`
}

var taskService *service.TaskService

func SetTaskService(s *service.TaskService) {
	taskService = s
}

type CreateRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
type UpdateRequest struct {
	Completed bool `json:"completed"`
}

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
func Tasks(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserIDKey).(string)

	switch r.Method {
	case http.MethodPost:
		var req CreateRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		if req.Title == "" || req.Description == "" {
			http.Error(w, "Title && Description is required", http.StatusBadRequest)
			return
		}
		task, err := taskService.CreateTask(req.Title, sql.NullString{String: req.Description, Valid: true}, userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		writeJSON(w, http.StatusCreated, task)
	case http.MethodGet:
		tasks, err := taskService.GetAllTasks(userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		writeJSON(w, http.StatusOK, tasks)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
func TaskById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := strings.TrimPrefix(r.URL.Path, "/tasks/")
	if id == "" {
		http.Error(w, "Task id required", http.StatusBadRequest)
		return
	}
	method := r.Method
	switch method {
	case http.MethodGet:
		task, err := taskService.GetTaskById(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		writeJSON(w, http.StatusOK, task)
	case http.MethodPatch:
		var req UpdateRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		err := taskService.UpdateTask(id, req.Completed)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		writeJSON(w, http.StatusOK, "Task Updated")
	case http.MethodDelete:
		err := taskService.DeleteTask(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		writeJSON(w, http.StatusOK, "Task deleted")
	default:
		fmt.Fprintln(w, "Method not allowed")
	}
}
