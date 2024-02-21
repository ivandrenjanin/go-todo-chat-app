package projectstore

import (
	"context"

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

func (ps ProjectStorage) FindProjectsByOwnerId(ctx context.Context, id int) ([]pg.Project, error,
) {
	projects, err := ps.store.Pg.ProjectsByUserId(ctx, id)
	if err != nil {
		return nil, err
	}

	return projects, nil
}
