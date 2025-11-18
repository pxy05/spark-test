package handlers

import (
	"backend/internal/database"
	"net/http"
)

func Middleware(w http.ResponseWriter, r *http.Request, db *database.Database) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	switch r.Method {
	case http.MethodGet:
		ToDoListHandler(w, r, db)
	case http.MethodPost:
		ToDoHandler(w, r, db)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)

	}

	// Your code here
}