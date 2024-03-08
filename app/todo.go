package app

import "context"

type TodoState struct {
	ID        int
	Name      string
	ItemOrder int
}

type Todo struct {
	ID          int
	Name        string
	Description string
	ItemOrder   int
}

type ToDoCollection struct {
	TodoState
	Todos []Todo
}

type TodoCollectionMap = map[string]ToDoCollection

type TodoStore interface {
	ToDosByProjectId(ctx context.Context, id int) (TodoCollectionMap, error)
}

type ToDoService struct {
	store TodoStore
}

func NewTodoService(store TodoStore) ToDoService {
	return ToDoService{
		store: store,
	}
}
