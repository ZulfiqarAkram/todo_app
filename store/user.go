package store

import (
	"fmt"
	"github.com/mzulfiqar10p/todo_app/model"
)

type UserStore interface {
	GetUserByEmailAddress(emailAddress string) (*model.User, error)
	AddUser(newUser model.User) error
	GetUserByEmailAndPassword(emailAddress string, password string) (*model.User, error)
}

func (Store *DBStore) GetUserByEmailAddress(emailAddress string) (*model.User, error) {
	user := &model.User{}
	if err := Store.DB.Where("email_address = ?", emailAddress).Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
func (Store *DBStore) AddUser(newUser model.User) error {
	if err := Store.DB.Create(&newUser).Error; err != nil {
		return err
	}
	fmt.Println("After New user added: Done")
	return nil
}
func (Store *DBStore) GetUserByEmailAndPassword(emailAddress string, password string) (*model.User, error) {
	var user = &model.User{}
	if err := Store.DB.Where("email_address = ? AND password = ?", emailAddress, password).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
