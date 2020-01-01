package main

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"log"
	"net/http"
	"todo_app/auth"
	"todo_app/middleware"
	"todo_app/routes"
	st "todo_app/store"
)

func main() {
	// Register JWT
	auth.CreateJWTManager()

	//Seed DB
	st.UploadMockData()

	//Create Routes
	r := mux.NewRouter()
	routes.Initialize(r)

	//Middleware
	n := negroni.New()
	n.UseHandler(middleware.NewAuthorization(r))

	log.Fatal(http.ListenAndServe(":8080", n))
}
