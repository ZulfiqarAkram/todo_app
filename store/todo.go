package store

import (
	"fmt"
	"github.com/mzulfiqar10p/todo_app/model"
)

type TodoStore interface {
	GetTodoItems() ([]model.TodoItem, error)
	AddTodo(newTodo model.TodoItem) error
	DeleteTodo(id int, userID int) error
	UpdateTodo(id int, userID int, todoToBeUpdate model.TodoItem) (*model.TodoItem, error)
	GetTodoItemsByUserID(userID int) ([]model.TodoItem, error)
	GetTodoItemByID(ID int) (*model.TodoItem, error)
}

func (Store *DBStore) GetTodoItems() ([]model.TodoItem, error) {
	var todoItems []model.TodoItem
	if err := Store.DB.Find(&todoItems).Error; err != nil {
		return nil, err
	}
	return todoItems, nil
}
func (Store *DBStore) AddTodo(newTodo model.TodoItem) error {
	if err := Store.DB.Create(&newTodo).Error; err != nil {
		return err
	}
	fmt.Println("AFTER ADDED : Done")
	return nil
}
func (Store *DBStore) DeleteTodo(id int) error {
	var todoItem model.TodoItem
	if err := Store.DB.Where("id = ?", id).Delete(&todoItem).Error; err != nil {
		return err
	}
	fmt.Println("AFTER REMOVED : Done")
	return nil
}
func (Store *DBStore) UpdateTodo(ID uint, todoToBeUpdate model.TodoItem) (*model.TodoItem, error) {
	todoItem := &model.TodoItem{}
	if err := Store.DB.Model(&todoItem).Where("id = ?", ID).Update("text", todoToBeUpdate.Text).Error; err != nil {
		return nil, err
	}
	fmt.Println("AFTER UPDATED : Done")
	return todoItem, nil
}
func (Store *DBStore) GetTodoItemsByUserID(userID uint) ([]model.TodoItem, error) {
	var todoItems []model.TodoItem
	if err := Store.DB.Where("user_id = ?", userID).Find(&todoItems).Error; err != nil {
		return nil, err
	}
	return todoItems, nil
}
func (Store *DBStore) GetTodoItemByID(ID int) (*model.TodoItem, error) {
	todoItem := &model.TodoItem{}
	if err := Store.DB.First(&todoItem, ID).Error; err != nil {
		return nil, err
	}
	return todoItem, nil
}
