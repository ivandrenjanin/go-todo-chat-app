package app

import "golang.org/x/net/context"

type UserService struct {
	store Store
}

func NewUserService(store Store) UserService {
	return UserService{
		store: store,
	}
}

func (us UserService) FindById(ctx context.Context, id int) (User, error) {
	u, err := us.store.FindById(ctx, id)
	return u, err
}
