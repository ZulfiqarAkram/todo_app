package api

import (
	"github.com/gorilla/mux"
	mw "todo_app/api/middleware"
	"todo_app/auth"
	store "todo_app/store"
	"todo_app/validator"
)

type API struct {
	MainRouter       *mux.Router
	Store            *store.DBStore
	JWTManager       *auth.JwtAuth
	Router           *Router
	ValidatorManager *validator.Validate
	Authentication   *mw.Authentication
}

func New() *API {
	api := &API{
		MainRouter:       mux.NewRouter(),
		Store:            store.New(),
		JWTManager:       auth.CreateJWTManager(),
		Router:           &Router{},
		ValidatorManager: validator.New(),
	}
	api.setupRoutes()
	return api
}

func (api *API) Initialize() {
	api.Store.SeedDatabase()
	api.Authentication = mw.New(api.MainRouter, api.JWTManager)
}
