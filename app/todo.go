package app

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type TodoState struct {
	ID        int
	Name      string
	ItemOrder int
	ProjectID int
}

type Todo struct {
	ID          int
	PublicID    uuid.UUID
	Name        string
	Description string
	ItemOrder   int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	StateID     int
}

type ToDoCollection struct {
	TodoState
	Todos []Todo
}

type TodoCollectionMap = map[string]ToDoCollection

type TodoStore interface {
	// ToDosByProjectId(ctx context.Context, id int) (TodoCollectionMap, error)
	ToDosByStateId(ctx context.Context, id int) ([]Todo, error)
	TodoStatesByProjectId(ctx context.Context, id int) ([]TodoState, error)
}

type ToDoService struct {
	store TodoStore
}

func NewTodoService(store TodoStore) ToDoService {
	return ToDoService{
		store: store,
	}
}

func (s ToDoService) FindTodosByProjectId(
	ctx context.Context,
	id int,
) ([]TodoState, TodoCollectionMap, error) {
	ts, err := s.store.TodoStatesByProjectId(ctx, id)
	if err != nil {
		return nil, nil, err
	}

	m := make(TodoCollectionMap)

	for _, v := range ts {
		t, err := s.store.ToDosByStateId(ctx, v.ID)
		if err != nil {
			return nil, nil, err
		}
		m[v.Name] = ToDoCollection{
			TodoState: v,
			Todos:     t,
		}

	}

	return ts, m, nil
}
