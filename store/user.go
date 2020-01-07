package store

import (
	"fmt"
	"todo_app/model"
)

type DBStore struct {
	userTbl []model.User
	todoTbl []model.TodoItem
}

func NewStore() *DBStore {
	return &DBStore{
		userTbl: []model.User{},
		todoTbl: []model.TodoItem{},
	}
}

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

func (db *DBStore) SeedDatabase() {
	db.userTbl = []model.User{
		{
			ID:           1,
			Username:     "aliahmed",
			Password:     "123456",
			EmailAddress: "aliahmed@gmail.com",
		},
		{
			ID:           2,
			Username:     "muhammad",
			Password:     "786786",
			EmailAddress: "muhammad@gmail.com",
		},
	}

	db.todoTbl = []model.TodoItem{
		{
			ID:     1,
			Text:   "Going for cricket",
			UserID: 1,
		}, {
			ID:     2,
			Text:   "Pickup my dad",
			UserID: 2,
		}, {
			ID:     3,
			Text:   "Going for lunch",
			UserID: 1,
		}, {
			ID:     4,
			Text:   "Get HD from Daraz",
			UserID: 2,
		},
	}
}
