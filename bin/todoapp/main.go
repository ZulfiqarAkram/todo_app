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
	auth.CreateJWTManager()
	st.UploadMockData()
	r := mux.NewRouter()
	routes.Initialize(r)

	n := negroni.New()
	n.UseHandler(middleware.NewAuthorization(r))

	log.Fatal(http.ListenAndServe(":8080", n))
}
