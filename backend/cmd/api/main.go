package main

import (
	"backend/internal/database"
	"backend/internal/handlers"
	"fmt"
	"net/http"
)

const PORT string = ":8080"

var db = database.NewDatabase()



func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request)  {
		handlers.Middleware(w, r, db)
	})

	fmt.Println("Server starting on port", PORT)
	http.ListenAndServe(PORT, mux)
}




