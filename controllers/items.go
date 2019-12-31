package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"todo_app/store"
	"todo_app/types"
)

func AddItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	token := r.Header.Get("token")
	payload, err, isValid := IsValidToken(token)
	if isValid {
		var usr types.User
		usr = usr.ConvertToStruct(payload)
		var newTodo types.TodoItem
		err := json.NewDecoder(r.Body).Decode(&newTodo)
		if err != nil {
			fmt.Println(err)
		}
		newTodo.UserID = usr.ID
		store.AddTodo(newTodo)
		err1 := json.NewEncoder(w).Encode(newTodo)
		if err1 != nil {
			fmt.Println(err1)
		}
	} else {
		err := json.NewEncoder(w).Encode(err)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	token := r.Header.Get("token")
	payload, err, isValid := IsValidToken(token)
	if isValid {
		var usr types.User
		usr = usr.ConvertToStruct(payload)
		id, err := strconv.ParseInt(params["id"], 16, 64)
		if err != nil {
			fmt.Println(err)
		}
		store.DeleteTodo(int(id), usr.ID)

		err1 := json.NewEncoder(w).Encode(store.GetTodoItems())
		if err1 != nil {
			fmt.Println(err1)
		}
	} else {
		err := json.NewEncoder(w).Encode(err)
		if err != nil {
			fmt.Println(err)
		}
	}

}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	token := r.Header.Get("token")
	payload, err, isValid := IsValidToken(token)
	if isValid {
		var usr types.User
		usr = usr.ConvertToStruct(payload)
		var todoToBeUpdate types.TodoItem
		id, err := strconv.ParseInt(params["id"], 16, 64)
		if err != nil {
			fmt.Println(err)
		}
		err1 := json.NewDecoder(r.Body).Decode(&todoToBeUpdate)
		if err1 != nil {
			fmt.Println(err1)
		}
		store.UpdateTodo(int(id), usr.ID, todoToBeUpdate)
		err2 := json.NewEncoder(w).Encode(todoToBeUpdate)
		if err2 != nil {
			fmt.Println(err2)
		}
	} else {
		err := json.NewEncoder(w).Encode(err)
		if err != nil {
			fmt.Println(err)
		}
	}

}

func DisplayItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	token := r.Header.Get("token")
	payload, err, isValid := IsValidToken(token)
	if isValid {
		var usr types.User
		usr = usr.ConvertToStruct(payload)
		todoItems := store.GetTodoItemsByUserID(usr.ID)
		err := json.NewEncoder(w).Encode(todoItems)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		err := json.NewEncoder(w).Encode(err)
		if err != nil {
			fmt.Println(err)
		}
	}
}
