package api

import (
	"encoding/json"
	"github.com/darahayes/go-boom"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"todo_app/model"
	"todo_app/util"
)

func (api *API) InitTodo() {
	api.Router.Todo.HandleFunc("", api.DisplayItems).Methods(http.MethodGet)
	api.Router.Todo.HandleFunc("", api.AddItem).Methods(http.MethodPost)
	api.Router.Todo.HandleFunc("/{id}", api.UpdateItem).Methods(http.MethodPut)
	api.Router.Todo.HandleFunc("/{id}", api.DeleteItem).Methods(http.MethodDelete)
}

func (api *API) AddItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	usr, err := model.GetUserFromContext(r)
	if err != nil {
		boom.Internal(w, err.Error())
		return
	}
	if err := api.ValidatorManager.ValidateStruct(usr); err != nil {
		boom.Internal(w, err.Error())
		return
	}

	var newTodo model.TodoItem
	newTodo.UserID = usr.ID
	err = json.NewDecoder(r.Body).Decode(&newTodo)
	if err != nil {
		boom.Internal(w, err.Error())
		return
	}
	api.Store.AddTodo(newTodo)
	err = util.JsonResponse(w, 200, "New todo item has been added.")
	if err != nil {
		boom.Internal(w, err.Error())
		return
	}
}
func (api *API) DisplayItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	usr, err := model.GetUserFromContext(r)
	if err != nil {
		boom.Internal(w, err.Error())
		return
	}
	if err := api.ValidatorManager.ValidateStruct(usr); err != nil {
		boom.Internal(w, err.Error())
		return
	}
	todoItems := api.Store.GetTodoItemsByUserID(usr.ID)
	err = json.NewEncoder(w).Encode(todoItems)
	if err != nil {
		boom.Internal(w, err.Error())
		return
	}
}
func (api *API) UpdateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	usr, err := model.GetUserFromContext(r)
	if err != nil {
		boom.Internal(w, err.Error())
		return
	}
	if err := api.ValidatorManager.ValidateStruct(usr); err != nil {
		boom.Internal(w, err.Error())
		return
	}
	var todoToBeUpdate model.TodoItem
	id, err := strconv.ParseInt(params["id"], 16, 64)
	if err != nil {
		boom.Internal(w, err.Error())
		return
	}
	err = json.NewDecoder(r.Body).Decode(&todoToBeUpdate)
	if err != nil {
		boom.Internal(w, err.Error())
		return
	}
	api.Store.UpdateTodo(int(id), usr.ID, todoToBeUpdate)
	todoItem := api.Store.GetTodoItemByID(int(id))
	err = json.NewEncoder(w).Encode(todoItem)
	if err != nil {
		boom.Internal(w, err.Error())
		return
	}
}
func (api *API) DeleteItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	usr, err := model.GetUserFromContext(r)
	if err != nil {
		boom.Internal(w, err.Error())
		return
	}
	if err := api.ValidatorManager.ValidateStruct(usr); err != nil {
		boom.Internal(w, err.Error())
		return
	}
	id, err := strconv.ParseInt(params["id"], 16, 64)
	if err != nil {
		boom.Internal(w, err.Error())
		return
	}
	api.Store.DeleteTodo(int(id), usr.ID)

	err = util.JsonResponse(w, 200, "Todo item has been deleted.")
	if err != nil {
		boom.Internal(w, err.Error())
		return
	}
}
