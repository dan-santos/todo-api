package handlers

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/dan-santos/todo-api/models"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo

	err := json.NewDecoder(r.Body).Decode(&todo); if err != nil {
		slog.Error("failed to decode request body")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := models.Insert(todo)

	var resp map[string]any

	if err != nil {
		resp = map[string]any{
			"Error": true,
			"Message": fmt.Sprintf("failed on insert given todo: %v", err),
		}
	} else {
		resp = map[string]any{
			"Error": false,
			"Message": fmt.Sprintf("todo successfully inserted: %d", id),
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}