package services

import "github.com/ivandrenjanin/go-chat-app/db"

type ToDoService struct {
	*db.Database
}

func NewTodo(db *db.Database) ToDoService {
	return ToDoService{db}
}
