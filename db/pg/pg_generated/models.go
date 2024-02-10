// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package pg

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Project struct {
	ID          int32     `json:"id"`
	PublicID    uuid.UUID `json:"publicId"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	OwnerID     int32     `json:"ownerId"`
}

type ProjectAssignment struct {
	ProjectID      int32 `json:"projectId"`
	UserID         int32 `json:"userId"`
	ProjectOwnerID int32 `json:"projectOwnerId"`
}

type Todo struct {
	ID          int32        `json:"id"`
	PublicID    uuid.UUID    `json:"publicId"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	CreatedAt   time.Time    `json:"createdAt"`
	UpdatedAt   time.Time    `json:"updatedAt"`
	DeletedAt   sql.NullTime `json:"deletedAt"`
	ProjectID   int32        `json:"projectId"`
}

type TodoAssignment struct {
	TodoID int32 `json:"todoId"`
	UserID int32 `json:"userId"`
}

type User struct {
	ID        int32        `json:"id"`
	FirstName string       `json:"firstName"`
	LastName  string       `json:"lastName"`
	Email     string       `json:"email"`
	Password  string       `json:"password"`
	CreatedAt time.Time    `json:"createdAt"`
	UpdatedAt time.Time    `json:"updatedAt"`
	DeletedAt sql.NullTime `json:"deletedAt"`
}