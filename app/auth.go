package app

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"github.com/ivandrenjanin/go-chat-app/cfg"
)

type AuthService struct {
	jwtCfg *cfg.JwtConfig
}

func NewAuthService(config *cfg.Config) AuthService {
	return AuthService{
		jwtCfg: &config.JwtConfig,
	}
}

func (a AuthService) HashPassword(pw string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pw), 14)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func (a AuthService) CompareString(pw, hashedPw string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPw), []byte(pw))
	return err == nil
}

type CustomClaims struct {
	UserID int
	jwt.RegisteredClaims
}

func (a AuthService) SignToken(userId int) (string, error) {
	claims := CustomClaims{
		userId,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "App",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(a.jwtCfg.Secret)
	if err != nil {
		return "", err
	}

	return ss, nil
}

func (a AuthService) ValidateToken(tok string) bool {
	t, err := jwt.ParseWithClaims(
		tok,
		&CustomClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return a.jwtCfg.Secret, nil
		},
	)
	if err != nil {
		return false
	} else if _, ok := t.Claims.(*CustomClaims); ok {
		return true
	}
	return false
}
