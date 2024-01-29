package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/dan-santos/todo-api/models"
)

func List(w http.ResponseWriter, r *http.Request) {
	todos, err := models.GetAll(); if err != nil {
		slog.Error("failed to list all todos")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}