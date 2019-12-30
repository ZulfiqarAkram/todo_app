package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	d "todo_app/data"
	"todo_app/types"
	hf "todo_app/hp_func"
)

func AddItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	token := r.Header.Get("token")
	payload, err, isValid := isValidToken(token)
	if isValid {
		usr := hf.ConvertMapToStruct(payload)
		var newTodo types.TodoItem
		json.NewDecoder(r.Body).Decode(&newTodo)
		newTodo.ID = len(d.TodoDB) + 1
		newTodo.UserID=usr.ID
		d.TodoDB = append(d.TodoDB, newTodo)
		fmt.Println("AFTER ADDED : ", d.TodoDB)
		json.NewEncoder(w).Encode(newTodo)
	} else {
		json.NewEncoder(w).Encode(err)
	}
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	token := r.Header.Get("token")
	payload, err, isValid := isValidToken(token)
	if isValid {
		usr := hf.ConvertMapToStruct(payload)
		for index, item := range d.TodoDB {
			id, err := strconv.ParseInt(params["id"], 16, 64)
			if err != nil {
				fmt.Println(err)
			}
			if item.ID == int(id) && usr.ID == item.UserID {
				d.TodoDB = append(d.TodoDB[:index], d.TodoDB[index+1:]...)
				fmt.Println("AFTER REMOVED : ", d.TodoDB)
				break
			}
		}
		json.NewEncoder(w).Encode(d.TodoDB)
	} else {
		json.NewEncoder(w).Encode(err)
	}

}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	token := r.Header.Get("token")
	payload, err, isValid := isValidToken(token)
	if isValid {
		usr := hf.ConvertMapToStruct(payload)
		var todoToBeUpdate types.TodoItem
		for index, item := range d.TodoDB {
			id, err := strconv.ParseInt(params["id"], 16, 64)
			if err != nil {
				fmt.Println(err)
			}
			if item.ID == int(id) && item.UserID == usr.ID {
				d.TodoDB = append(d.TodoDB[:index], d.TodoDB[index+1:]...)
				json.NewDecoder(r.Body).Decode(&todoToBeUpdate)
				todoToBeUpdate.ID = int(id)
				d.TodoDB = append(d.TodoDB, todoToBeUpdate)
				fmt.Println("AFTER UPDATED : ", d.TodoDB)
				break
			}
		}
		json.NewEncoder(w).Encode(todoToBeUpdate)
	} else {
		json.NewEncoder(w).Encode(err)
	}

}

func DisplayItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	token := r.Header.Get("token")
	payload, err, isValid := isValidToken(token)
	fmt.Println(payload," | ", err," | ", isValid)
	if isValid {
		usr := hf.ConvertMapToStruct(payload)
		json.NewEncoder(w).Encode(hf.GetUserTodoItems(usr.ID))
	} else {
		json.NewEncoder(w).Encode(err)
	}
}
