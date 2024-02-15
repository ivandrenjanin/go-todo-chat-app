package services

import (
	"golang.org/x/net/context"
)

type UserService struct {
	storage     *UserStorage
	authService *AuthService
}

func NewUserService(s *UserStorage, as *AuthService) UserService {
	return UserService{
		storage:     s,
		authService: as,
	}
}

func (s UserService) RegisterUser(
	ctx context.Context,
	fn string,
	ln string,
	em string,
	pw string,
) (int, error) {
	hashedPw, err := s.authService.HashPassword(pw)
	if err != nil {
		return 0, err
	}

	id, err := s.storage.Save(ctx, fn, ln, em, hashedPw)
	if err != nil {
		return 0, err
	}

	return id, nil
}
