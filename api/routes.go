package api

import "github.com/gorilla/mux"

type Router struct {
	APIRoot *mux.Router
	User    *mux.Router
	Todo    *mux.Router
}

func (api *API) setupRoutes() {
	api.Router.APIRoot = api.MainRouter.PathPrefix("/api/v1").Subrouter()
	api.Router.User = api.Router.APIRoot.PathPrefix("/user").Subrouter()
	api.Router.Todo = api.Router.APIRoot.PathPrefix("/todo").Subrouter()

	api.InitUser()
	api.InitTodo()
}
