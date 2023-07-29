package main

import (
	"net/http"

	"api/db"
	"api/models"
	"api/routes"

	"github.com/gorilla/mux"
)

func main() {

	db.DBConection()
	// db.DB.Migrator().DropTable(models.User{})
	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})

	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)
	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users", routes.PostUsersHandler).Methods("POST")
	r.HandleFunc("/users/{id}", routes.DeleteUsersHandler).Methods("DELETE")

	http.ListenAndServe(":3000", r)
}
