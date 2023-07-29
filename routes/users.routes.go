package routes

import "net/http"

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get users"))

}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get user"))

}

func PostUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("post user"))

}

func DeleteUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete user"))

}
