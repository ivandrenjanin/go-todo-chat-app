package app

import (
	"errors"

	"golang.org/x/net/context"

	"github.com/ivandrenjanin/go-chat-app/storage"
)

type UserService struct {
	storage     *storage.UserStorage
	authService *AuthService
}

func NewUserService(s *storage.UserStorage, as *AuthService) UserService {
	return UserService{
		storage:     s,
		authService: as,
	}
}

// TODO: Register should go to AuthService
func (s UserService) RegisterUser(
	ctx context.Context,
	fn string,
	ln string,
	em string,
	pw string,
) (string, error) {
	hashedPw, err := s.authService.HashPassword(pw)
	if err != nil {
		return "", err
	}

	id, err := s.storage.Save(ctx, fn, ln, em, hashedPw)
	if err != nil {
		return "", err
	}

	token, err := s.authService.SignToken(id)
	if err != nil {
		return "", err
	}
	return token, nil
}

// TODO: Login should go to AuthService
func (s UserService) Login(ctx context.Context, em string, pw string) (string, error) {
	u, err := s.storage.Pg.UserByEmail(ctx, em)
	if err != nil {
		return "", errors.New("Can not find user")
	}

	ok := s.authService.CompareString(pw, u.Password)
	if !ok {
		return "", errors.New("Invalid password")
	}

	token, err := s.authService.SignToken(u.ID)
	if err != nil {
		return "", errors.New("Can not sign token")
	}

	return token, nil
}
