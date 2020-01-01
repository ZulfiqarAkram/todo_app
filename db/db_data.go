package db

import (
	"todo_app/types"
)

var UserDB []types.User

var TodoDB []types.TodoItem

func SeedData() {
	UserDB = []types.User{
		{
			ID:           1,
			Username:     "Ali",
			Password:     "123456",
			EmailAddress: "ali@gmail.com",
		},
		{
			ID:           2,
			Username:     "Muhammad",
			Password:     "786786",
			EmailAddress: "muhammad@gmail.com",
		},
	}

	TodoDB = []types.TodoItem{
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
