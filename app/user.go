package app

import "golang.org/x/net/context"

type UserService struct {
	store UserStore
}

type UserStore interface {
	FindById(ctx context.Context, id int) (User, error)
}

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Password  string
}

func NewUserService(store UserStore) UserService {
	return UserService{
		store: store,
	}
}

func (us UserService) FindById(ctx context.Context, id int) (User, error) {
	u, err := us.store.FindById(ctx, id)
	return u, err
}
