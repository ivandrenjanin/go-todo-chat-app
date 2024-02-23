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

type ProjectAssignment struct {
	ProjectID      int
	UserID         int
	ProjectOwnerID int
}

type ProjectCollection struct {
	Project
	ProjectAssignment
}

type ProjectStore interface {
	ProjectById(ctx context.Context, id int) (Project, error)
	ProjectsByUserId(ctx context.Context, id int) ([]ProjectCollection, error,
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

func (ps ProjectService) FindProjectsByUserId(
	ctx context.Context,
	userId int,
) ([]ProjectCollection, error) {
	return ps.store.ProjectsByUserId(ctx, userId)
}
