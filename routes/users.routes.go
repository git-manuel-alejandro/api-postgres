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

func PostUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	createdUser := db.DB.Create(&user)
	err := createdUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		m := Message{err.Error()}

		json.NewEncoder(w).Encode(m)
	} else {

		json.NewEncoder(w).Encode(&user)
	}

}

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var users []models.User
	db.DB.Find(&users)
	if len(users) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		m := Message{"Not users registered"}
		json.NewEncoder(w).Encode(&m)
		return

	}

	json.NewEncoder(w).Encode(&users)

}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var user models.User
	params := mux.Vars(r)
	db.DB.First(&user, params["id"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		m := Message{"User not found"}
		json.NewEncoder(w).Encode(&m)

	} else {
		db.DB.Model(&user).Association("Tasks").Find(&user.Tasks)
		json.NewEncoder(w).Encode(user)

	}

}

func DeleteUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var user models.User
	params := mux.Vars(r)
	db.DB.First(&user, params["id"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)

		m := Message{"user not found"}
		json.NewEncoder(w).Encode(m)

	} else {
		// db.DB.Delete(&user) // marca deleted
		db.DB.Unscoped().Delete(&user) // borra de la tabla
		w.WriteHeader(http.StatusOK)
		m := Message{"User id " + params["id"] + " was deleted"}
		json.NewEncoder(w).Encode(m)

	}

}
