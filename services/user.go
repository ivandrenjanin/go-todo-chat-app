package services

import "github.com/ivandrenjanin/go-chat-app/db"

type UserService struct {
	*db.Database
}

func NewUser(db *db.Database) UserService {
	return UserService{db}
}
