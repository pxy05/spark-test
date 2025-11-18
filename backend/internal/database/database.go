package database

import (
	"backend/internal/types"
	"errors"
	"strings"
	"sync"
)

type Database struct {
	Data types.ToDoList
	mu   sync.RWMutex
}

func (db *Database) AddTodo(toDoItem types.ToDo) (types.ToDo, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	toDoItem.Title = strings.TrimSpace(toDoItem.Title)
	toDoItem.Description = strings.TrimSpace(toDoItem.Description)

	if ( toDoItem.Title == "" || toDoItem.Description == "" ) {
		return types.ToDo{}, errors.New("Error: Empty Title or Empty Description")
	}

	db.Data = append(db.Data, toDoItem)
	return db.Data[len(db.Data)-1], nil
}

func (db *Database) GetAll() types.ToDoList {
	db.mu.RLock()
	defer db.mu.RUnlock()

	return db.Data // copy by value no race condition
}

func NewDatabase() *Database {
	return &Database{Data: make(types.ToDoList, 0)}
}