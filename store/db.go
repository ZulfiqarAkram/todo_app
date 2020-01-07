package store

import "todo_app/model"

type DBStore struct {
	userTbl []model.User
	todoTbl []model.TodoItem
}

//NewStore constructor
func New() *DBStore {
	return &DBStore{
		userTbl: []model.User{},
		todoTbl: []model.TodoItem{},
	}
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
