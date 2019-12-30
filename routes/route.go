package routes

import (
	"github.com/gorilla/mux"
	"net/http"
	"todo_app/controllers"
)


func Initialize(r *mux.Router) {
	api := r.PathPrefix("/api").Subrouter()

	//Authentication apis
	api.HandleFunc("/login", controllers.Login).Methods(http.MethodPost)
	api.HandleFunc("/register", controllers.Register).Methods(http.MethodPost)

	//_Todo basic app apis
	api.HandleFunc("/todo", controllers.DisplayItems).Methods(http.MethodGet)
	api.HandleFunc("/todo", controllers.AddItem).Methods(http.MethodPost)
	api.HandleFunc("/todo/{id}", controllers.UpdateItem).Methods(http.MethodPut)
	api.HandleFunc("/todo/{id}", controllers.DeleteItem).Methods(http.MethodDelete)
}