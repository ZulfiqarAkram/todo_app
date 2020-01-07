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

func (db *DBStore) GetUserByEmailAddress(emailAddress string) model.User {
	var user model.User
	for _, usr := range db.userTbl {
		if usr.EmailAddress == emailAddress {
			user = usr
			break
		}
	}
	return user
}
func (db *DBStore) AddUser(newUser model.User) {
	newUser.ID = len(db.userTbl) + 1
	db.userTbl = append(db.userTbl, newUser)
	fmt.Println("After New user added: ", db.userTbl)
}
func (db *DBStore) GetUserByEmailAndPassword(emailAddress string, password string) model.User {
	var selUsr model.User
	for _, usr := range db.userTbl {
		if usr.EmailAddress == emailAddress && usr.Password == password {
			selUsr = usr
			break
		}
	}
	return selUsr
}
