package handlers

import (
	"encoding/json"
	"github.com/ferhangedik/simple-go-web-app/db"
	"github.com/gorilla/mux"
	"net/http"
)

// GetTasks handles the GET request to retrieve all tasks
func GetTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []db.Task
	db.DB.Find(&tasks)
	json.NewEncoder(w).Encode(tasks)
}

// GetTaskByID handles the GET request to retrieve a task by its ID
func GetTaskByID(w http.ResponseWriter, r *http.Request) {
	var task db.Task
	id := mux.Vars(r)["id"]
	db.DB.First(&task, id)
	json.NewEncoder(w).Encode(task)
}

// CreateTask handles the POST request to create a new task
func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task db.Task
	_ = json.NewDecoder(r.Body).Decode(&task)
	db.DB.Create(&task)
	json.NewEncoder(w).Encode(task)
}

// UpdateTask handles the PUT request to update a task by its ID
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	var task db.Task
	id := mux.Vars(r)["id"]
	_ = json.NewDecoder(r.Body).Decode(&task)
	db.DB.Model(&task).Where("id = ?", id).Updates(&task)
	json.NewEncoder(w).Encode(task)
}

// DeleteTask handles the DELETE request to delete a task by its ID
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	db.DB.Delete(&db.Task{}, id)
	json.NewEncoder(w).Encode(map[string]string{"message": "Task deleted"})
}
