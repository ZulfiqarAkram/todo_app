package types

import (
	"encoding/json"
	"fmt"
	d "todo_app/db"
)

type User struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	EmailAddress string `json:"email_address"`
}
type TodoItem struct {
	ID     int    `json:"id"`
	Text   string `json:"text"`
	UserID int    `json:"user_id"`
}

func (u *User) GetUserByEmailAndPassword(emailAddress string, password string) User {
	var selUsr User
	for _, usr := range d.UserDB {
		if usr.EmailAddress == emailAddress && usr.Password == password {
			selUsr = usr
			break
		}
	}
	return selUsr
}
func (u *User) ConvertToStruct(payload map[string]interface{}) User {
	jsonBody, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
	}
	usr := User{}
	err1 := json.Unmarshal(jsonBody, &usr)
	if err1 != nil {
		fmt.Println(err1)
	}
	return usr
}
func (ti *TodoItem) GetTodoItemsByUserID(userID int) []TodoItem {
	var fTodoItems []TodoItem

	for _, todo := range d.TodoDB {
		if todo.UserID == userID {
			fTodoItems = append(fTodoItems, todo)
		}
	}
	return fTodoItems
}
