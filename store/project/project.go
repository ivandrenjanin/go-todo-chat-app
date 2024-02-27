package projectstore

import (
	"context"

	"github.com/google/uuid"

	"github.com/ivandrenjanin/go-chat-app/app"
	"github.com/ivandrenjanin/go-chat-app/db"
	pg "github.com/ivandrenjanin/go-chat-app/db/pg/pg_generated"
)

type ProjectStorage struct {
	store *db.Database
}

func New(s *db.Database) ProjectStorage {
	return ProjectStorage{
		store: s,
	}
}

func (s ProjectStorage) ProjectsByUserId(
	ctx context.Context,
	id int,
) ([]app.ProjectCollection, error,
) {
	p, err := s.store.Pg.ProjectsByUserId(ctx, id)
	if err != nil {
		return nil, err
	}

	projects := make([]app.ProjectCollection, 0, cap(p))
	for _, v := range p {
		projects = append(projects, app.ProjectCollection{
			Project:           v.Project.ConvertToProject(),
			ProjectAssignment: v.ProjectAssignment.ConvertToProjectAssignment(),
		})
	}

	return projects, nil
}

func (s ProjectStorage) ProjectById(ctx context.Context, id string) (app.Project, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return app.Project{}, err
	}

	p, err := s.store.Pg.ProjectById(ctx, uid)
	if err != nil {
		return app.Project{}, err
	}

	return p.ConvertToProject(), nil
}

func (s ProjectStorage) DeleteProject(ctx context.Context, id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	return s.store.Pg.DeleteProject(ctx, uid)
}

func (s ProjectStorage) Save(
	ctx context.Context,
	u app.User,
	name, description string,
) (app.ProjectCollection, error) {
	tx, err := s.store.Db.Begin()
	if err != nil {
		return app.ProjectCollection{}, err
	}
	defer tx.Rollback()

	qtx := s.store.Pg.WithTx(tx)

	pArgs := pg.InsertProjectParams{
		Name:        name,
		Description: description,
		OwnerID:     u.ID,
	}

	p, err := qtx.InsertProject(ctx, pArgs)
	if err != nil {
		return app.ProjectCollection{}, err
	}

	paArgs := pg.InsertProjectAssignmentParams{
		ProjectID:      p.ID,
		UserID:         u.ID,
		ProjectOwnerID: u.ID,
	}

	pa, err := qtx.InsertProjectAssignment(ctx, paArgs)
	if err != nil {
		return app.ProjectCollection{}, err
	}

	err = tx.Commit()
	if err != nil {
		return app.ProjectCollection{}, err
	}

	return app.ProjectCollection{
		Project:           p.ConvertToProject(),
		ProjectAssignment: pa.ConvertToProjectAssignment(),
	}, nil
}
