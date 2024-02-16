package app

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
)

type AuthService struct {
	jwtCfg  jwtConfig
	storage userStorage
}

type jwtConfig interface {
	GetJwtSecret() []byte
}

type userStorage interface {
	Save(context.Context, string, string, string, string) (int, error)
	FindUserByEmail(context.Context, string) (struct {
		ID       int
		Password string
	}, error)
}

func NewAuthService(jwtConfig jwtConfig, storage userStorage) AuthService {
	return AuthService{
		jwtCfg:  jwtConfig,
		storage: storage,
	}
}

func (s AuthService) hashPassword(pw string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pw), 14)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func (s AuthService) compareString(pw, hashedPw string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPw), []byte(pw))
	return err == nil
}

type CustomClaims struct {
	UserID int
	jwt.RegisteredClaims
}

func (s AuthService) signToken(userId int) (string, error) {
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
	ss, err := token.SignedString(s.jwtCfg.GetJwtSecret())
	if err != nil {
		return "", err
	}

	return ss, nil
}

func (s AuthService) ValidateToken(tok string) bool {
	t, err := jwt.ParseWithClaims(
		tok,
		&CustomClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return s.jwtCfg.GetJwtSecret(), nil
		},
	)
	if err != nil {
		return false
	} else if _, ok := t.Claims.(*CustomClaims); ok {
		return true
	}
	return false
}

func (s AuthService) Register(
	ctx context.Context,
	fn string,
	ln string,
	em string,
	pw string,
) (string, error) {
	hashedPw, err := s.hashPassword(pw)
	if err != nil {
		return "", err
	}

	id, err := s.storage.Save(ctx, fn, ln, em, hashedPw)
	if err != nil {
		return "", err
	}

	token, err := s.signToken(id)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s AuthService) Login(ctx context.Context, em string, pw string) (string, error) {
	u, err := s.storage.FindUserByEmail(ctx, em)
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
