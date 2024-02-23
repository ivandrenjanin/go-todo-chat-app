package projectstore

import (
	"context"

	"github.com/ivandrenjanin/go-chat-app/app"
	"github.com/ivandrenjanin/go-chat-app/db"
)

type ProjectStorage struct {
	store *db.Database
}

func New(s *db.Database) ProjectStorage {
	return ProjectStorage{
		store: s,
	}
}

func (s ProjectStorage) ProjectsByUserId(ctx context.Context, id int) ([]app.Project, error,
) {
	p, err := s.store.Pg.ProjectsByUserId(ctx, id)
	if err != nil {
		return nil, err
	}

	projects := make([]app.Project, len(p), cap(p))

	return projects, nil
}

func (s ProjectStorage) ProjectById(ctx context.Context, id int) (app.Project, error) {
	p, err := s.store.Pg.ProjectById(ctx, id)
	if err != nil {
		return app.Project{}, err
	}

	return p.ConvertToProject(), nil
}
