package controllers

import (
	"encoding/json"
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if usr == (types.User{}) {
		err := json.NewEncoder(w).Encode("user object not found.")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	var newTodo types.TodoItem
	newTodo.UserID = usr.ID
	err1 := json.NewDecoder(r.Body).Decode(&newTodo)
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusInternalServerError)
		return
	}
	store.AddTodo(newTodo)
	err2 := json.NewEncoder(w).Encode(newTodo)
	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusInternalServerError)
		return
	}
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	usr, err := getUserFromContext(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if usr != (types.User{}) {
		id, err := strconv.ParseInt(params["id"], 16, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		store.DeleteTodo(int(id), usr.ID)

		err1 := json.NewEncoder(w).Encode(store.GetTodoItems())
		if err1 != nil {
			http.Error(w, err1.Error(), http.StatusInternalServerError)
		}
	} else {
		err := json.NewEncoder(w).Encode("user object not found.")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	usr, err := getUserFromContext(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if usr != (types.User{}) {
		var todoToBeUpdate types.TodoItem
		id, err := strconv.ParseInt(params["id"], 16, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		err1 := json.NewDecoder(r.Body).Decode(&todoToBeUpdate)
		if err1 != nil {
			http.Error(w, err1.Error(), http.StatusInternalServerError)
		}
		store.UpdateTodo(int(id), usr.ID, todoToBeUpdate)
		err2 := json.NewEncoder(w).Encode(todoToBeUpdate)
		if err2 != nil {
			http.Error(w, err2.Error(), http.StatusInternalServerError)
		}
	} else {
		err := json.NewEncoder(w).Encode("user object not found.")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

}

func DisplayItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	usr, err := getUserFromContext(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if usr != (types.User{}) {
		todoItems := store.GetTodoItemsByUserID(usr.ID)
		err := json.NewEncoder(w).Encode(todoItems)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		err := json.NewEncoder(w).Encode("user object not found.")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

}

func getUserFromContext(r *http.Request) (types.User, error) {
	payload := r.Context().Value(middleware.AuthenticatedUserKey)
	var usr types.User
	usr, err := usr.ConvertToStruct(payload)
	return usr, err
}
