package main

import (
	"github.com/urfave/negroni"
	"log"
	"net/http"
	"todo_app/api"
)

func main() {
	//Create API and Initialize Routes, DB Store, Validator & JWT
	a, err  := api.New()
	if err != nil {
		log.Fatal(err)
	}

	a.Initialize()

	//Middleware
	n := negroni.New()
	n.UseHandler(a.Authentication)

	log.Fatal(http.ListenAndServe(":8080", n))
}
