package api

import (
	"github.com/gorilla/mux"
	mw "todo_app/api/middleware"
	"todo_app/auth"
	st "todo_app/store"
	"todo_app/validator"
)

type API struct {
	MainRouter     *mux.Router
	MyStore        *st.DBStore
	JWTManager     *auth.JwtAuth
	Router         *Router
	MyValidator    *validator.VStruct
	Authentication *mw.Authentication
}

func NewAPI() *API {
	api := &API{
		MainRouter:  mux.NewRouter(),
		MyStore:     st.NewStore(),
		JWTManager:  auth.CreateJWTManager(),
		Router:      &Router{},
		MyValidator: validator.New(),
	}
	api.setupRoutes()
	return api
}

func (api *API) Initialize() {
	api.MyStore.SeedDatabase()
	api.Authentication = mw.New(api.MainRouter, api.JWTManager)
}
