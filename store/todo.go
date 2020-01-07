package store

import (
	"fmt"
	"todo_app/model"
)

type TodoStore interface {
	GetTodoItems() []model.TodoItem
	AddTodo(newTodo model.TodoItem)
	DeleteTodo(id int, userID int)
	UpdateTodo(id int, userID int, todoToBeUpdate model.TodoItem)
	GetTodoItemsByUserID(userID int) []model.TodoItem
	GetTodoItemByID(ID int) model.TodoItem
}

func (db *DBStore) GetTodoItems() []model.TodoItem {
	return db.todoTbl
}
func (db *DBStore) AddTodo(newTodo model.TodoItem) {
	newTodo.ID = len(db.todoTbl) + 1
	db.todoTbl = append(db.todoTbl, newTodo)
	fmt.Println("AFTER ADDED : ", db.todoTbl)
}
func (db *DBStore) DeleteTodo(id int, userID int) {
	for index, item := range db.todoTbl {
		if item.ID == id && userID == item.UserID {
			db.todoTbl = append(db.todoTbl[:index], db.todoTbl[index+1:]...)
			fmt.Println("AFTER REMOVED : ", db.todoTbl)
			break
		}
	}
}
func (db *DBStore) UpdateTodo(id int, userID int, todoToBeUpdate model.TodoItem) {
	for index, item := range db.todoTbl {
		if item.ID == id && item.UserID == userID {
			db.todoTbl = append(db.todoTbl[:index], db.todoTbl[index+1:]...)
			todoToBeUpdate.ID = id
			todoToBeUpdate.UserID = userID
			db.todoTbl = append(db.todoTbl, todoToBeUpdate)
			fmt.Println("AFTER UPDATED : ", db.todoTbl)
			break
		}
	}
}
func (db *DBStore) GetTodoItemsByUserID(userID int) []model.TodoItem {
	var fTodoItems []model.TodoItem
	for _, todo := range db.todoTbl {
		if todo.UserID == userID {
			fTodoItems = append(fTodoItems, todo)
		}
	}
	return fTodoItems
}
func (db *DBStore) GetTodoItemByID(ID int) model.TodoItem {
	var todoItem model.TodoItem
	for _, todo := range db.todoTbl {
		if todo.ID == ID {
			todoItem = todo
			break
		}
	}
	return todoItem
}
