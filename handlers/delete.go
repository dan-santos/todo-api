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

func Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id")); if err != nil {
		slog.Error("failed to decode request url param")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	rows, err := models.Delete(int64(id));

	var resp map[string]any

	if err != nil || rows != 1 {
		resp = map[string]any{
			"Error": true,
			"Message": fmt.Sprintf("failed on remove given todo: %v", err),
		}
	} else {
		resp = map[string]any{
			"Error": false,
			"Message": fmt.Sprintf("todo successfully removed: %d", id),
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}