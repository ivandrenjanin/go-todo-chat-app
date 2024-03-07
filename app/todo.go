package app

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

type ToDoService struct{}

func NewTodoService() ToDoService {
	return ToDoService{}
}
