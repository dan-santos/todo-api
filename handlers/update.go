package handlers

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/dan-santos/todo-api/models"
	"github.com/go-chi/chi/v5"
)

func Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id")); if err != nil {
		slog.Error("failed to decode request url param")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var todo models.Todo

	err = json.NewDecoder(r.Body).Decode(&todo); if err != nil {
		slog.Error("failed to decode request body")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	_, err = models.Update(int64(id), todo);

	var resp map[string]any

	if err != nil {
		resp = map[string]any{
			"Error": true,
			"Message": fmt.Sprintf("failed on update given todo: %v", err),
		}
	} else {
		resp = map[string]any{
			"Error": false,
			"Message": fmt.Sprintf("todo successfully updated: %d", id),
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}