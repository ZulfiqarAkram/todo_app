package controllers

import (
	"encoding/json"
	"github.com/darahayes/go-boom"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"todo_app/middleware"
	"todo_app/store"
	"todo_app/types"
)

func AddItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	usr, err := getUserFromContext(r)
	if err != nil {
		boom.Internal(w, err.Error())
		return
	}
	if err := myValidator.ValidateStruct(usr); err != nil {
		boom.Internal(w, err.Error())
		return
	}

	var newTodo types.TodoItem
	newTodo.UserID = usr.ID
	err = json.NewDecoder(r.Body).Decode(&newTodo)
	if err != nil {
		boom.Internal(w, err.Error())
		return
	}
	store.AddTodo(newTodo)
	err = JsonResponse(w, 200, "New todo item has been added.")
	if err != nil {
		boom.Internal(w, err.Error())
		return
	}
}
func DisplayItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	usr, err := getUserFromContext(r)
	if err != nil {
		boom.Internal(w, err.Error())
		return
	}
	if err := myValidator.ValidateStruct(usr); err != nil {
		boom.Internal(w, err.Error())
		return
	}
	todoItems := store.GetTodoItemsByUserID(usr.ID)
	err = json.NewEncoder(w).Encode(todoItems)
	if err != nil {
		boom.Internal(w, err.Error())
		return
	}
}
func UpdateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	usr, err := getUserFromContext(r)
	if err != nil {
		boom.Internal(w, err.Error())
		return
	}
	if err := myValidator.ValidateStruct(usr); err != nil {
		boom.Internal(w, err.Error())
		return
	}
	var todoToBeUpdate types.TodoItem
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
	store.UpdateTodo(int(id), usr.ID, todoToBeUpdate)
	todoItem := store.GetTodoItemByID(int(id))
	err = json.NewEncoder(w).Encode(todoItem)
	if err != nil {
		boom.Internal(w, err.Error())
		return
	}
}
func DeleteItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	usr, err := getUserFromContext(r)
	if err != nil {
		boom.Internal(w, err.Error())
		return
	}
	if err := myValidator.ValidateStruct(usr); err != nil {
		boom.Internal(w, err.Error())
		return
	}
	id, err := strconv.ParseInt(params["id"], 16, 64)
	if err != nil {
		boom.Internal(w, err.Error())
		return
	}
	store.DeleteTodo(int(id), usr.ID)

	err = JsonResponse(w, 200, "Todo item has been deleted.")
	if err != nil {
		boom.Internal(w, err.Error())
		return
	}
}

//Helper functions
func getUserFromContext(r *http.Request) (types.User, error) {
	payload := r.Context().Value(middleware.AuthenticatedUserKey)
	var usr types.User
	if err := usr.ConvertToStruct(payload); err != nil {
		return usr, err
	}
	return usr, nil
}
