package app

import (
	"context"
	"errors"
	"time"

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
	CreatedAt   time.Time
	UpdatedAt   time.Time
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
	ProjectById(ctx context.Context, id string) (Project, error)
	ProjectsByUserId(ctx context.Context, id int) ([]ProjectCollection, error)
	DeleteProject(ctx context.Context, id string) error
	Save(
		ctx context.Context,
		u User,
		name, description string,
	) (ProjectCollection, error)
}

func NewProjectService(store ProjectStore) ProjectService {
	return ProjectService{
		store: store,
	}
}

func (ps ProjectService) FindProjectById(ctx context.Context, id string) (Project, error) {
	return ps.store.ProjectById(ctx, id)
}

func (ps ProjectService) FindProjectsByUserId(
	ctx context.Context,
	userId int,
) ([]ProjectCollection, error) {
	return ps.store.ProjectsByUserId(ctx, userId)
}

func (ps ProjectService) RemoveProject(ctx context.Context, u User, id string) error {
	p, err := ps.FindProjectById(ctx, id)
	if err != nil {
		return err
	}

	if p.OwnerID != u.ID {
		return errors.New("Forbidden Operation")
	}

	return ps.store.DeleteProject(ctx, id)
}

func (ps ProjectService) CreateProject(
	ctx context.Context,
	u User,
	name, description string,
) (ProjectCollection, error) {
	return ps.store.Save(ctx, u, name, description)
}
