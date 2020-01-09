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

func New() (*API, error) {
	newStore, err := store.New()
	if err != nil {
		return nil, err
	}
	newValidator, err := validator.New()
	if err != nil {
		return nil, err
	}

	api := &API{
		Store:            newStore,
		MainRouter:       mux.NewRouter(),
		JWTManager:       auth.CreateJWTManager(),
		Router:           &Router{},
		ValidatorManager: newValidator,
	}
	api.setupRoutes()
	return api, nil
}

func (api *API) Initialize() {
	api.Store.SeedDatabase()
	api.Authentication = mw.New(api.MainRouter, api.JWTManager)
}
