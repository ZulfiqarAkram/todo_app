package hp_func

import (
	"encoding/json"
	"fmt"
	d "todo_app/data"
	"todo_app/types"
)

func ConvertMapToStruct(payload map[string]interface{}) types.User {
	jsonBody, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
	}
	usr := types.User{}
	err1 := json.Unmarshal(jsonBody, &usr)
	if err1 != nil {
		fmt.Println(err1)
	}
	return usr
}

func GetUserTodoItems(userID int) []types.TodoItem {
	var fTodoItems []types.TodoItem

	for _, todo := range d.TodoDB {
		if todo.UserID == userID {
			fTodoItems = append(fTodoItems, todo)
		}
	}
	return fTodoItems
}

func IsDuplicateUser(emailAddress string) bool {
	for _, usr := range d.UserDB {
		if usr.EmailAddress == emailAddress {
			return true
		}
	}
	return false
}

func IsAuthenticUser(emailAddress string, password string) bool {
	for _, usr := range d.UserDB {
		if usr.EmailAddress == emailAddress && usr.Password == password {
			return true
		}
	}
	return false
}

func GetUser(emailAddress string, password string) types.User {
	var selUsr types.User
	for _, usr := range d.UserDB {
		if usr.EmailAddress == emailAddress && usr.Password == password {
			selUsr = usr
			break
		}
	}
	return selUsr
}
