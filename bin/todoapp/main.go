package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"todo_app/auth"
	"todo_app/routes"
	st "todo_app/store"
)

func main() {
	auth.CreateJWTManager()
	st.UploadMockData()
	r := mux.NewRouter()
	routes.Initialize(r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
