package types

import (
	"encoding/json"
	"fmt"
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

func (u *User) ConvertToStruct(payload interface{}) User {
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
