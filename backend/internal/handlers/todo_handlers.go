package handlers

import (
	"backend/internal/database"
	"backend/pkg/types"
	"encoding/json"
	"net/http"
)

func ToDoHandler(w http.ResponseWriter, r *http.Request, db *database.Database) {
	w.Header().Set("Content-Type", "application/json")
	var toDoItem types.ToDo

	err := json.NewDecoder(r.Body).Decode(&toDoItem)

	if err != nil {
		http.Error(w, "Invalid Input", http.StatusBadRequest)
		return
	}

	createdToDo, err := db.AddTodo(toDoItem)

	if err != nil {
		http.Error(w, "Invalid Input", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(createdToDo)
}

func ToDoListHandler(w http.ResponseWriter, r *http.Request, db *database.Database) {
	w.Header().Set("Content-Type", "application/json")
	var dbCopy types.ToDoList = db.GetAll()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dbCopy)
}