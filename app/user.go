package app

import (
	"fmt"

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

	fmt.Printf("User Service - PW: %#v\n", hashedPw)

	id, err := s.storage.Save(ctx, fn, ln, em, hashedPw)
	if err != nil {
		return 0, err
	}

	fmt.Printf("User Service - User Save: %#v\n", id)

	return id, nil
}
