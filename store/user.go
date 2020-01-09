package store

import (
	"fmt"
	"todo_app/model"
)

type UserStore interface {
	GetUserByEmailAddress(emailAddress string) model.User
	AddUser(newUser model.User)
	GetUserByEmailAndPassword(emailAddress string, password string) model.User
}

func (Store *DBStore) GetUserByEmailAddress(emailAddress string) model.User {
	var user model.User
	if err := Store.DB.Where("email_address = ?", emailAddress).Find(&user).Error; err != nil {
		fmt.Println(err.Error())
	}
	return user
}
func (Store *DBStore) AddUser(newUser model.User) {
	if err := Store.DB.Create(&newUser).Error; err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("After New user added: Done")
}
func (Store *DBStore) GetUserByEmailAndPassword(emailAddress string, password string) model.User {
	var user model.User
	if err := Store.DB.Where("email_address = ? AND password = ?", emailAddress, password).First(&user).Error; err != nil {
		fmt.Println(err.Error())
	}
	return user
}
