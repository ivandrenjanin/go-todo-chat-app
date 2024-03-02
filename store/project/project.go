package projectstore

import (
	"context"
	"time"

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

func (s ProjectStorage) SaveProjectAssignment(
	ctx context.Context,
	p app.Project,
	u app.User,
) (app.ProjectAssignment, error) {
	paArgs := pg.InsertProjectAssignmentParams{
		ProjectID:      p.ID,
		UserID:         u.ID,
		ProjectOwnerID: p.OwnerID,
	}

	pa, err := s.store.Pg.InsertProjectAssignment(ctx, paArgs)
	if err != nil {
		return app.ProjectAssignment{}, err
	}

	return pa.ConvertToProjectAssignment(), nil
}

func (s ProjectStorage) SaveInvitation(
	ctx context.Context,
	p app.Project,
	email, token string,
) (app.ProjectInvitation, error) {
	args := pg.InsertProjectInvitationParams{
		ProjectID: p.ID,
		Email:     email,
		Token:     token,
		SentAt:    time.Now(),
		ExpiresAt: time.Now().Add(time.Duration(48 * time.Hour)),
	}

	i, err := s.store.Pg.InsertProjectInvitation(ctx, args)
	if err != nil {
		return app.ProjectInvitation{}, err
	}

	return i.ConvertToProjectInvitation(), err
}
