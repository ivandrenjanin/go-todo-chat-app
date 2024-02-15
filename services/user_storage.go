package services

import (
	"golang.org/x/net/context"

	"github.com/ivandrenjanin/go-chat-app/db"
)

type UserStorage struct {
	*db.Database
}

func NewUserStorage(db *db.Database) UserStorage {
	return UserStorage{
		db,
	}
}

func (s UserStorage) Save(ctx context.Context,
	fn string,
	ln string,
	em string,
	pw string,
) (int, error) {
	arg := struct {
		FirstName string
		LastName  string
		Email     string
		Password  string
	}{
		FirstName: fn,
		LastName:  ln,
		Email:     em,
		Password:  pw,
	}
	id, err := s.Pg.InsertUser(ctx, arg)
	if err != nil {
		return 0, nil
	}

	return id, nil
}
