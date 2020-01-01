package types

import (
	"encoding/json"
)

type User struct {
	ID           int    `json:"id" validate:"required"`
	Username     string `json:"username" validate:"required,min=5,max=15"`
	Password     string `json:"password" validate:"required,pwdMinLenSix"`
	EmailAddress string `json:"email_address" validate:"required,email"`
}

type LoginUser struct {
	Password     string `json:"password" validate:"required,pwdMinLenSix"`
	EmailAddress string `json:"email_address" validate:"required,email"`
}

type RegisterUser struct {
	Username     string `json:"username" validate:"required,min=5,max=15"`
	Password     string `json:"password" validate:"required,pwdMinLenSix"`
	EmailAddress string `json:"email_address" validate:"required,email"`
}

type TodoItem struct {
	ID     int    `json:"id"`
	Text   string `json:"text" validate:"required"`
	UserID int    `json:"user_id"`
}

func (u *User) ConvertToStruct(payload interface{}) error {
	jsonBody, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(jsonBody, &u); err != nil {
		return err
	}
	return nil
}
