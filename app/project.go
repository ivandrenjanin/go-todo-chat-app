package app

import (
	"context"

	"github.com/google/uuid"
)

type ProjectService struct {
	store ProjectStore
}

type Project struct {
	ID          int
	PublicID    uuid.UUID
	Name        string
	Description string
	OwnerID     int
}

type ProjectStore interface {
	ProjectById(ctx context.Context, id int) (Project, error)
	ProjectsByUserId(ctx context.Context, id int) ([]Project, error,
	)
}

func NewProjectService(store ProjectStore) ProjectService {
	return ProjectService{
		store: store,
	}
}

func (ps ProjectService) FindProjectById(ctx context.Context, id int) (Project, error) {
	return ps.store.ProjectById(ctx, id)
}

func (ps ProjectService) FindProjectsByUserId(ctx context.Context, userId int) ([]Project, error) {
	return ps.store.ProjectsByUserId(ctx, userId)
}
