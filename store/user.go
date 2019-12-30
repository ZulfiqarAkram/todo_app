package store

import (
	d "todo_app/db"
	"todo_app/types"
)

func GetUserByEmailAddress(emailAddress string) types.User {
	var user types.User
	for _, usr := range d.UserDB {
		if usr.EmailAddress == emailAddress {
			user = usr
			break
		}
	}
	return user
}
