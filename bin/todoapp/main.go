package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	d "todo_app/db"
	"todo_app/routes"
)

func main() {
	d.SeedData()
	r := mux.NewRouter()
	routes.Initialize(r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
