package main

import (
	"github.com/ferhangedik/simple-go-web-app/db"
	"github.com/ferhangedik/simple-go-web-app/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	// Initialize the database
	db.InitDatabase()

	// Create a new Gorilla Mux router
	router := mux.NewRouter()

	// Set up routes for tasks
	taskRouter := router.PathPrefix("/tasks").Subrouter()
	taskRouter.HandleFunc("/", handlers.GetTasks).Methods("GET")
	taskRouter.HandleFunc("/{id}", handlers.GetTaskByID).Methods("GET")
	taskRouter.HandleFunc("/", handlers.CreateTask).Methods("POST")
	taskRouter.HandleFunc("/{id}", handlers.UpdateTask).Methods("PUT")
	taskRouter.HandleFunc("/{id}", handlers.DeleteTask).Methods("DELETE")

	// Serve static files (e.g., for documentation)
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Run the server on port 8080
	log.Fatal(http.ListenAndServe(":8080", router))
}
