package services

import "github.com/ivandrenjanin/go-chat-app/db"

type ProjectService struct {
	*db.Database
}

func NewProject(db *db.Database) ProjectService {
	return ProjectService{db}
}
