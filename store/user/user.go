package userstore

import (
	"context"

	"github.com/ivandrenjanin/go-chat-app/app"
	"github.com/ivandrenjanin/go-chat-app/db"
)

type UserStorage struct {
	*db.Database
}

func New(db *db.Database) UserStorage {
	return UserStorage{
		db,
	}
}

func (s UserStorage) Save(ctx context.Context, u app.User) (int, error) {
	arg := struct {
		FirstName string
		LastName  string
		Email     string
		Password  string
	}{
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Password:  u.Password,
	}
	id, err := s.Pg.InsertUser(ctx, arg)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s UserStorage) FindByEmail(ctx context.Context, em string) (app.User, error,
) {
	u, err := s.Pg.UserByEmail(ctx, em)
	if err != nil {
		return app.User{}, err
	}

	return u.ConvertToUser(), nil
}

func (s UserStorage) FindById(ctx context.Context, id int) (app.User, error) {
	u, err := s.Pg.User(ctx, id)
	if err != nil {
		return app.User{}, err
	}

	return u.ConvertToUser(), nil
}
