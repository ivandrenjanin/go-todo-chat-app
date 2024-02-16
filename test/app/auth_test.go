package app_test

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/bcrypt"

	"github.com/ivandrenjanin/go-chat-app/app"
)

// Mock Dependency
type jwtConfig struct {
	*authServiceSuite
}

func (j jwtConfig) GetJwtSecret() []byte {
	j.getJwtSecretCallCounter++
	return []byte("test")
}

// Mock Dependency
type storage struct {
	*authServiceSuite
}

func (s storage) Save(ctx context.Context, fn, ln, em, pw string) (int, error) {
	s.storageSaveCallCounter++

	if strings.Contains(em, "invalid") {
		return 0, errors.New("error")
	}

	return 1, nil
}

func (s storage) FindUserByEmail(ctx context.Context, em string) (struct {
	ID       int
	Password string
}, error,
) {
	s.storageFindUserByEmailCallCounter++
	var res struct {
		ID       int
		Password string
	}
	if strings.Contains(em, "invalid") {
		return res, errors.New("error")
	}

	res.Password = s.validPassword
	res.ID = 1
	return res, nil
}

// Suite Setup
type authServiceSuite struct {
	suite.Suite

	as app.AuthService

	jwtConfig interface {
		GetJwtSecret() []byte
	}
	getJwtSecretCallCounter int

	storage interface {
		Save(context.Context, string, string, string, string) (int, error)
		FindUserByEmail(context.Context, string) (struct {
			ID       int
			Password string
		}, error)
	}
	storageSaveCallCounter            int
	storageFindUserByEmailCallCounter int
	validPassword                     string
}

func (suite *authServiceSuite) SetupSuite() {
	suite.jwtConfig = jwtConfig{
		suite,
	}
	suite.storage = storage{
		suite,
	}

	vp, _ := bcrypt.GenerateFromPassword([]byte("valid-password"), 14)
	suite.validPassword = string(vp)

	suite.as = app.NewAuthService(suite.jwtConfig, suite.storage)
}

func (suite *authServiceSuite) AfterTest() {
	suite.getJwtSecretCallCounter = 0
	suite.storageSaveCallCounter = 0
	suite.storageFindUserByEmailCallCounter = 0
}

// Suite Test Cases
func (suite *authServiceSuite) TestLoginShouldReturnErrorOnInvalidEmail() {
	// Arrange
	beforeCount := suite.storageFindUserByEmailCallCounter
	ctx := context.Background()

	// Act
	tok, err := suite.as.Login(ctx, "invalid-mail@email.com", "1234")

	// Assert
	assert.Equal(suite.T(), "", tok)
	assert.ErrorContains(suite.T(), err, "Can not find user")
	suite.NotEqual(suite.storageFindUserByEmailCallCounter, beforeCount)
}

func (suite *authServiceSuite) TestLoginShouldReturnErrorOnInvalidPassword() {
	// Arrange
	beforeCount := suite.storageFindUserByEmailCallCounter
	ctx := context.Background()

	// Act
	tok, err := suite.as.Login(ctx, "mail@email.com", "invalid-password")

	// Assert
	assert.Equal(suite.T(), "", tok)
	assert.ErrorContains(suite.T(), err, "Invalid password")
	suite.NotEqual(suite.storageFindUserByEmailCallCounter, beforeCount)
}

func (suite *authServiceSuite) TestLoginShouldReturnValidTokenOnValidPassword() {
	// Arrange
	beforeFindUserCount := suite.storageFindUserByEmailCallCounter
	beforeGetSecretCount := suite.getJwtSecretCallCounter
	ctx := context.Background()

	// Act
	tok, err := suite.as.Login(ctx, "mail@email.com", "valid-password")

	// Assert
	assert.NotEqual(suite.T(), "", tok)
	assert.Equal(suite.T(), err, nil)
	suite.NotEqual(suite.storageFindUserByEmailCallCounter, beforeFindUserCount)
	suite.NotEqual(suite.getJwtSecretCallCounter, beforeGetSecretCount)
}

func (suite *authServiceSuite) TestRegisterShouldReturnErrorOnLongPassword() {
	// Arrange
	beforeCount := suite.storageSaveCallCounter
	ctx := context.Background()
	pw := "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxA"

	// Act
	tok, err := suite.as.Register(ctx, "fn", "ln", "em", pw)

	// Assert
	assert.Equal(suite.T(), "", tok)
	assert.ErrorContains(suite.T(), err, "Password too long")
	suite.Equal(suite.storageSaveCallCounter, beforeCount)
}

func (suite *authServiceSuite) TestRegisterShouldReturnErrorOnFailedSave() {
	// Arrange
	beforeCount := suite.storageSaveCallCounter
	ctx := context.Background()

	// Act
	tok, err := suite.as.Register(ctx, "fn", "ln", "invalid-mail@email.com", "pw")

	// Assert
	assert.Equal(suite.T(), "", tok)
	assert.ErrorContains(suite.T(), err, "Could not save a new user to storage")
	suite.NotEqual(suite.storageSaveCallCounter, beforeCount)
}

func (suite *authServiceSuite) TestRegisterShouldReturnValidToken() {
	// Arrange
	beforeSaveCount := suite.storageSaveCallCounter
	beforeGetSecretCount := suite.getJwtSecretCallCounter
	ctx := context.Background()

	// Act
	tok, err := suite.as.Register(ctx, "fn", "ln", "em", "pw")

	// Assert
	assert.NotEqual(suite.T(), "", tok)
	assert.Equal(suite.T(), err, nil)
	suite.NotEqual(suite.storageSaveCallCounter, beforeSaveCount)
	suite.NotEqual(suite.getJwtSecretCallCounter, beforeGetSecretCount)
}

// Run our suite
func TestAuthServiceSuite(t *testing.T) {
	suite.Run(t, new(authServiceSuite))
}
