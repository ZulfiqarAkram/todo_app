package api

import (
	"github.com/gorilla/mux"
	"todo_app/auth"
	st "todo_app/store"
	"todo_app/validator"
)

type API struct {
	MainRouter  *mux.Router
	MyStore     *st.DBStore
	JWTManager  *auth.JwtAuth
	Router      *Router
	MyValidator *validator.VStruct
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
}

func (api *API) IsValidToken(token string) (map[string]interface{}, error) {
	payLoadResult, err := api.JWTManager.Decode(token)
	if err != nil {
		return payLoadResult, err
	}
	return payLoadResult, nil
}
