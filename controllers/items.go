package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	d "todo_app/db"
	"todo_app/types"
)

func AddItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	token := r.Header.Get("token")
	payload, err, isValid := isValidToken(token)
	if isValid {
		var usr types.User
		usr = usr.ConvertToStruct(payload)
		var newTodo types.TodoItem
		err := json.NewDecoder(r.Body).Decode(&newTodo)
		if err != nil {
			fmt.Println(err)
		}
		newTodo.ID = len(d.TodoDB) + 1
		newTodo.UserID = usr.ID
		d.TodoDB = append(d.TodoDB, newTodo)
		fmt.Println("AFTER ADDED : ", d.TodoDB)
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
	payload, err, isValid := isValidToken(token)
	if isValid {
		var usr types.User
		usr = usr.ConvertToStruct(payload)
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
		err := json.NewEncoder(w).Encode(d.TodoDB)
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

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	token := r.Header.Get("token")
	payload, err, isValid := isValidToken(token)
	if isValid {
		var usr types.User
		usr = usr.ConvertToStruct(payload)
		var todoToBeUpdate types.TodoItem
		for index, item := range d.TodoDB {
			id, err := strconv.ParseInt(params["id"], 16, 64)
			if err != nil {
				fmt.Println(err)
			}
			if item.ID == int(id) && item.UserID == usr.ID {
				d.TodoDB = append(d.TodoDB[:index], d.TodoDB[index+1:]...)
				err := json.NewDecoder(r.Body).Decode(&todoToBeUpdate)
				if err != nil {
					fmt.Println(err)
				}
				todoToBeUpdate.ID = int(id)
				d.TodoDB = append(d.TodoDB, todoToBeUpdate)
				fmt.Println("AFTER UPDATED : ", d.TodoDB)
				break
			}
		}
		err := json.NewEncoder(w).Encode(todoToBeUpdate)
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

func DisplayItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	token := r.Header.Get("token")
	payload, err, isValid := isValidToken(token)
	fmt.Println(payload, " | ", err, " | ", isValid)
	if isValid {
		var usr types.User
		usr = usr.ConvertToStruct(payload)
		var todoItem types.TodoItem
		todoItems := todoItem.GetTodoItemsByUserID(usr.ID)
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
