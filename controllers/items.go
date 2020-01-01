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
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := myValidator.ValidateStruct(usr); err != nil {
		boom.Internal(w, err.Error())
		//http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var newTodo types.TodoItem
	newTodo.UserID = usr.ID
	err = json.NewDecoder(r.Body).Decode(&newTodo)
	if err != nil {
		boom.Internal(w, err.Error())
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	store.AddTodo(newTodo)
	err = json.NewEncoder(w).Encode(newTodo)
	if err != nil {
		boom.Internal(w, err.Error())
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	usr, err := getUserFromContext(r)
	if err != nil {
		boom.Internal(w, err.Error())
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := myValidator.ValidateStruct(usr); err != nil {
		boom.Internal(w, err.Error())
		//http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := strconv.ParseInt(params["id"], 16, 64)
	if err != nil {
		boom.Internal(w, err.Error())
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	store.DeleteTodo(int(id), usr.ID)

	err = json.NewEncoder(w).Encode(store.GetTodoItems())
	if err != nil {
		boom.Internal(w, err.Error())
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	usr, err := getUserFromContext(r)
	if err != nil {
		boom.Internal(w, err.Error())
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := myValidator.ValidateStruct(usr); err != nil {
		boom.Internal(w, err.Error())
		//http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var todoToBeUpdate types.TodoItem
	id, err := strconv.ParseInt(params["id"], 16, 64)
	if err != nil {
		boom.Internal(w, err.Error())
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&todoToBeUpdate)
	if err != nil {
		boom.Internal(w, err.Error())
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	store.UpdateTodo(int(id), usr.ID, todoToBeUpdate)
	err = json.NewEncoder(w).Encode(todoToBeUpdate)
	if err != nil {
		boom.Internal(w, err.Error())
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func DisplayItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	usr, err := getUserFromContext(r)
	if err != nil {
		boom.Internal(w, err.Error())
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := myValidator.ValidateStruct(usr); err != nil {
		boom.Internal(w, err.Error())
		//http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	todoItems := store.GetTodoItemsByUserID(usr.ID)
	err = json.NewEncoder(w).Encode(todoItems)
	if err != nil {
		boom.Internal(w, err.Error())
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getUserFromContext(r *http.Request) (types.User, error) {
	payload := r.Context().Value(middleware.AuthenticatedUserKey)
	var usr types.User
	if err := usr.ConvertToStruct(payload); err != nil {
		return usr, err
	}
	return usr, nil
}
