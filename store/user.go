package store

import (
	"fmt"
	db "todo_app/db"
	"todo_app/types"
)

//type UserStore interface {
//	GetUserByEmailAddress(emailAddress string) types.User
//	AddUser(newUser types.User)
//	GetUserByEmailAndPassword(emailAddress string, password string) types.User
//}

//func ABC(x Userstore) {}

//type userStore struct{}

//NewStore NewStore
//func NewStore() UserStore {
//	u := &userStore{}
//	ABC(u)
//	return u
//}

func GetUserByEmailAddress(emailAddress string) types.User {
	var user types.User
	for _, usr := range db.UserDB {
		if usr.EmailAddress == emailAddress {
			user = usr
			break
		}
	}
	return user
}

func AddUser(newUser types.User) {
	newUser.ID = len(db.UserDB) + 1
	db.UserDB = append(db.UserDB, newUser)
	fmt.Println("After New user added: ", db.UserDB)
}

func UploadMockData() {
	db.SeedData()
}

func GetUserByEmailAndPassword(emailAddress string, password string) types.User {
	var selUsr types.User
	for _, usr := range db.UserDB {
		if usr.EmailAddress == emailAddress && usr.Password == password {
			selUsr = usr
			break
		}
	}
	return selUsr
}
