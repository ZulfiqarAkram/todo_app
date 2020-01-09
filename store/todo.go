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

func (Store *DBStore) GetTodoItems() []model.TodoItem {
	var todoItems []model.TodoItem
	if err := Store.DB.Find(&todoItems).Error; err != nil {
		fmt.Println(err.Error())
	}
	return todoItems
}
func (Store *DBStore) AddTodo(newTodo model.TodoItem) {
	if err := Store.DB.Create(&newTodo).Error; err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("AFTER ADDED : Done")

}
func (Store *DBStore) DeleteTodo(id int) {
	var todoItem model.TodoItem
	if err := Store.DB.Where("id = ?", id).Delete(&todoItem).Error; err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("AFTER REMOVED : Done")
}
func (Store *DBStore) UpdateTodo(ID uint, todoToBeUpdate model.TodoItem) model.TodoItem {
	var todoItem model.TodoItem
	if err := Store.DB.Model(&todoItem).Where("id = ?", ID).Update("text", todoToBeUpdate.Text).Error; err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("AFTER UPDATED : Done")
	return todoItem
}
func (Store *DBStore) GetTodoItemsByUserID(userID uint) []model.TodoItem {
	var todoItems []model.TodoItem
	if err := Store.DB.Where("user_id = ?", userID).Find(&todoItems).Error; err != nil {
		fmt.Println(err.Error())
	}
	return todoItems
}
func (Store *DBStore) GetTodoItemByID(ID int) model.TodoItem {
	var todoItem model.TodoItem
	if err := Store.DB.First(&todoItem, ID).Error; err != nil {
		fmt.Println(err.Error())
	}
	return todoItem
}
