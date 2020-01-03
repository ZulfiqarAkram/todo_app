package store

import (
	"fmt"
	"todo_app/db"
	"todo_app/types"
)

func GetTodoItems() []types.TodoItem {
	return db.TodoDB
}
func AddTodo(newTodo types.TodoItem) {
	newTodo.ID = len(db.TodoDB) + 1
	db.TodoDB = append(db.TodoDB, newTodo)
	fmt.Println("AFTER ADDED : ", db.TodoDB)
}
func DeleteTodo(id int, userID int) {
	for index, item := range db.TodoDB {
		if item.ID == int(id) && userID == item.UserID {
			db.TodoDB = append(db.TodoDB[:index], db.TodoDB[index+1:]...)
			fmt.Println("AFTER REMOVED : ", db.TodoDB)
			break
		}
	}
}
func UpdateTodo(id int, userID int, todoToBeUpdate types.TodoItem) {
	for index, item := range db.TodoDB {
		if item.ID == int(id) && item.UserID == userID {
			db.TodoDB = append(db.TodoDB[:index], db.TodoDB[index+1:]...)
			todoToBeUpdate.ID = int(id)
			todoToBeUpdate.UserID = userID
			db.TodoDB = append(db.TodoDB, todoToBeUpdate)
			fmt.Println("AFTER UPDATED : ", db.TodoDB)
			break
		}
	}
}
func GetTodoItemsByUserID(userID int) []types.TodoItem {
	var fTodoItems []types.TodoItem

	for _, todo := range db.TodoDB {
		if todo.UserID == userID {
			fTodoItems = append(fTodoItems, todo)
		}
	}
	return fTodoItems
}
func GetTodoItemByID(ID int) types.TodoItem {
	var todoItem types.TodoItem
	for _, todo := range db.TodoDB {
		if todo.ID == ID {
			todoItem = todo
			break
		}
	}
	return todoItem
}
