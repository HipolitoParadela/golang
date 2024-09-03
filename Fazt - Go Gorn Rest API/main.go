package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hipolitoparadela/go-gorn-restapi/db"
	"github.com/hipolitoparadela/go-gorn-restapi/models"
	"github.com/hipolitoparadela/go-gorn-restapi/routes"
)

func main() {

	// DataBase - Postgress
	db.DBConection()

	db.DB.AutoMigrate(models.Tasks{})
	db.DB.AutoMigrate(models.User{})
	//

	// Routes
	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)
	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")

	r.HandleFunc("/", routes.HomeHandler)
	r.HandleFunc("/tasks", routes.GetTasksHandler).Methods("GET")
	r.HandleFunc("/tasks", routes.CreateTaskHandler).Methods("POST")
	r.HandleFunc("/tasks/{id}", routes.GetTaskHandler).Methods("GET")
	r.HandleFunc("/tasks/{id}", routes.DeleteTaskHandler).Methods("DELETE")

	http.ListenAndServe(":3000", r)
	//
}
