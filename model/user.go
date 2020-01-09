package model

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"net/http"
	"todo_app/api/middleware"
)

type User struct {
	gorm.Model
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

func GetUserFromContext(r *http.Request) (User, error) {
	payload := r.Context().Value(middleware.AuthenticatedUserKey)
	var usr User
	if err := usr.ConvertToStruct(payload); err != nil {
		return usr, err
	}
	return usr, nil
}
