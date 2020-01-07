package model

type TodoItem struct {
	ID     int    `json:"id"`
	Text   string `json:"text" validate:"required"`
	UserID int    `json:"user_id"`
}
