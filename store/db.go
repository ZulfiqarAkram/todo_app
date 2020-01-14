package store

import (
	"bytes"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/mzulfiqar10p/todo_app/model"
	"os/exec"
)

type DBStore struct {
	DB *gorm.DB
}

//NewStore constructor
func New() (*DBStore, error) {
	dbStore := &DBStore{}
	db, err := dbStore.getDB()
	if err != nil {
		return nil, err
	}
	dbStore.DB = db
	return dbStore, nil
}
func (Store *DBStore) getDB() (*gorm.DB, error) {
	//return gorm.Open("sqlite3", "todo_app.db")
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=todoapp_db1 password=123456")
	if err != nil {
		err = createDatabase()
		if err != nil {
			return nil, err
		} else {
			db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=todoapp_db1 password=123456")
			if err != nil {
				return nil, err
			}
			return db, nil
		}
	}
	return db, nil
}
func createDatabase() error {
	cmd := exec.Command("createdb", "-p", "5432", "-h", "127.0.0.1", "-U", "postgres", "-e", "todoapp_db1", "PGPASSWORD=PGPASSWORD")
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
func (Store *DBStore) SeedDatabase() {
	Store.DB.LogMode(true)
	Store.DB.AutoMigrate(&model.User{}, &model.TodoItem{})

	//User
	var usr = model.User{
		Username:     "ali",
		Password:     "1234",
		EmailAddress: "ali@gmail.com",
	}
	var usrInDB []model.User
	Store.DB.Where("email_address = ?", usr.EmailAddress).Find(&usrInDB)
	if len(usrInDB) == 0 {
		Store.DB.NewRecord(&usr)
		Store.DB.Create(&usr)
		fmt.Println("User has been Seed.")

	}

	//_Todo
	var todoItem = model.TodoItem{
		Text:   "Going for cricket",
		UserID: usr.ID,
	}
	var todoItemsInDB []model.TodoItem
	Store.DB.Where("text = ?", todoItem.Text).Find(&todoItemsInDB)
	if len(todoItemsInDB) == 0 {
		Store.DB.NewRecord(&todoItem)
		Store.DB.Create(&todoItem)
		fmt.Println("Todo has been Seed.")
	}
}
