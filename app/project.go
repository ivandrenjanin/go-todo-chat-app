package app

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

	"github.com/ivandrenjanin/go-chat-app/views/templates"
)

type Project struct {
	ID          int
	PublicID    uuid.UUID
	Name        string
	Description string
	OwnerID     int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ProjectAssignment struct {
	ProjectID      int
	UserID         int
	ProjectOwnerID int
}

type ProjectInvitation struct {
	ProjectID       int
	Email           string
	SentAt          time.Time
	ExpiresAt       time.Time
	InvitationToken string
}

type ProjectCollection struct {
	Project
	ProjectAssignment
}

type ProjectStore interface {
	ProjectById(ctx context.Context, id string) (Project, error)
	ProjectsByUserId(ctx context.Context, id int) ([]ProjectCollection, error)
	DeleteProject(ctx context.Context, id string) error
	Save(
		ctx context.Context,
		u User,
		name, description string,
	) (ProjectCollection, error)
	SaveInvitation(ctx context.Context, p Project, email, token string) (ProjectInvitation, error)
}

type Mailer interface {
	Send(to, subject, body string) error
}

type ProjectService struct {
	store  ProjectStore
	mailer Mailer
}

func NewProjectService(store ProjectStore, mailer Mailer) ProjectService {
	return ProjectService{
		store:  store,
		mailer: mailer,
	}
}

func (s ProjectService) FindProjectById(ctx context.Context, id string) (Project, error) {
	return s.store.ProjectById(ctx, id)
}

func (s ProjectService) FindProjectsByUserId(
	ctx context.Context,
	userId int,
) ([]ProjectCollection, error) {
	return s.store.ProjectsByUserId(ctx, userId)
}

func (s ProjectService) RemoveProject(ctx context.Context, u User, id string) error {
	p, err := s.FindProjectById(ctx, id)
	if err != nil {
		return err
	}

	if p.OwnerID != u.ID {
		return errors.New("Forbidden Operation")
	}

	return s.store.DeleteProject(ctx, id)
}

func (s ProjectService) CreateProject(
	ctx context.Context,
	u User,
	name, description string,
) (ProjectCollection, error) {
	return s.store.Save(ctx, u, name, description)
}

type ProjectCustomClaims struct {
	Email string
	jwt.RegisteredClaims
}

func (s ProjectService) signToken(email, secret string) (string, error) {
	claims := ProjectCustomClaims{
		email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "App",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return ss, nil
}

func (s ProjectService) ValidateToken(tok, secret string) (*ProjectCustomClaims, bool) {
	t, err := jwt.ParseWithClaims(
		tok,
		&CustomClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		},
	)
	if err != nil {
		return nil, false
	} else if c, ok := t.Claims.(*ProjectCustomClaims); ok {
		return c, true
	}
	return nil, false
}

func (s ProjectService) CreateInvitation(
	ctx context.Context,
	publicId string,
	email string,
) (ProjectInvitation, error) {
	p, err := s.FindProjectById(ctx, publicId)
	if err != nil {
		return ProjectInvitation{}, err
	}

	t, err := s.signToken(email, publicId)
	if err != nil {
		return ProjectInvitation{}, err
	}

	pi, err := s.store.SaveInvitation(ctx, p, email, t)
	if err != nil {
		return ProjectInvitation{}, err
	}

	link := fmt.Sprintf("https://%s:%d/api/project/invitation/?token=%s", "localhost", 3000, t)

	var buf bytes.Buffer
	templates.AssignUser(link, p.Name).Render(ctx, &buf)
	if err != nil {
		return ProjectInvitation{}, err
	}

	err = s.mailer.Send(
		email,
		fmt.Sprintf("You are invited to join a project %s", p.Name),
		buf.String(),
	)
	if err != nil {
		return pi, err
	}

	return pi, nil
}
