package app

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type IdentityService struct {
	cfg   Config
	store IdentityStore
}

type IdentityStore interface {
	Save(ctx context.Context, u User) (int, error)
	FindByEmail(ctx context.Context, em string) (User, error)
}

type Config interface {
	GetJwtSecret() []byte
}

func NewIdentityService(cfg Config, store IdentityStore) IdentityService {
	return IdentityService{
		cfg:   cfg,
		store: store,
	}
}

func (s IdentityService) hashPassword(pw string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pw), 14)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func (s IdentityService) compareString(pw, hashedPw string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPw), []byte(pw))
	return err == nil
}

type CustomClaims struct {
	UserID int
	jwt.RegisteredClaims
}

func (s IdentityService) signToken(userId int) (string, error) {
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
	ss, err := token.SignedString(s.cfg.GetJwtSecret())
	if err != nil {
		return "", err
	}

	return ss, nil
}

func (s IdentityService) ValidateToken(tok string) (*CustomClaims, bool) {
	t, err := jwt.ParseWithClaims(
		tok,
		&CustomClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return s.cfg.GetJwtSecret(), nil
		},
	)
	if err != nil {
		return nil, false
	} else if c, ok := t.Claims.(*CustomClaims); ok {
		return c, true
	}
	return nil, false
}

func (s IdentityService) Register(
	ctx context.Context,
	fn string,
	ln string,
	em string,
	pw string,
) (string, error) {
	hashedPw, err := s.hashPassword(pw)
	if err != nil {
		return "", errors.New("Password too long")
	}

	id, err := s.store.Save(ctx, User{FirstName: fn, LastName: ln, Email: em, Password: hashedPw})
	if err != nil {
		return "", errors.New("Could not save a new user to storage")
	}

	token, err := s.signToken(id)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s IdentityService) Login(ctx context.Context, em string, pw string) (string, error) {
	u, err := s.store.FindByEmail(ctx, em)
	if err != nil {
		return "", errors.New("Can not find user")
	}

	ok := s.compareString(pw, u.Password)
	if !ok {
		return "", errors.New("Invalid password")
	}

	token, err := s.signToken(u.ID)
	if err != nil {
		return "", errors.New("Can not sign token")
	}

	return token, nil
}
