package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/dan-santos/todo-api/models"
	"github.com/go-chi/chi/v5"
)

func Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id")); if err != nil {
		slog.Error("failed to decode request url param")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	todo, err := models.Get(int64(id));

	if err != nil {
		slog.Error("failed to get todo with given id")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	} 

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}