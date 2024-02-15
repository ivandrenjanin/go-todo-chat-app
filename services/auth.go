package services

import "golang.org/x/crypto/bcrypt"

type AuthService struct{}

func NewAuthService() AuthService {
	return AuthService{}
}

func (a AuthService) HashPassword(pw string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pw), 14)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func CompareString(pw, hashedPw string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPw), []byte(pw))
	return err == nil
}
