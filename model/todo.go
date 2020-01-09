package model

import "github.com/jinzhu/gorm"

type TodoItem struct {
	gorm.Model
	Text   string `json:"text" validate:"required"`
	UserID uint   `json:"user_id"`
	User   User   `json:"user_id",gorm:"association_foreignkey:UserID"`
}
