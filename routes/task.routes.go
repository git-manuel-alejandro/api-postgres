package routes

import (
	"api/db"
	"api/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Message struct {
	Msg string
}

func CreateTasksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)
	createdTask := db.DB.Create(&task)
	err := createdTask.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(&task)

}

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var task models.Task
	params := mux.Vars(r)
	db.DB.First(&task, params["id"])

	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		m := Message{"Task not found"}
		json.NewEncoder(w).Encode(m)
		return

	}

	json.NewEncoder(w).Encode(&task)

}

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var tasks []models.Task
	db.DB.Find(&tasks)
	json.NewEncoder(w).Encode(&tasks)

}

func DeleteTasksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	var task models.Task
	db.DB.First(&task, params["id"])

	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		m := Message{"Task not found"}
		json.NewEncoder(w).Encode(m)
		return
	}

	db.DB.Unscoped().Delete(&task)
	w.WriteHeader(http.StatusOK)
	m := Message{"Task id " + params["id"] + " was deleted"}
	json.NewEncoder(w).Encode(m)

}
