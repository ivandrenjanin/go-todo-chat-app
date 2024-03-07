package todostore

import (
	"context"

	"github.com/ivandrenjanin/go-chat-app/app"
	"github.com/ivandrenjanin/go-chat-app/db"
)

type TodoStorage struct {
	store *db.Database
}

func New(s *db.Database) TodoStorage {
	return TodoStorage{
		store: s,
	}
}

func (s TodoStorage) ToDosByProjectId(ctx context.Context, id int) (app.TodoCollectionMap, error) {
	var result app.TodoCollectionMap

	rows, err := s.store.Pg.ToDosByProjectId(ctx, id)
	if err != nil {
		return result, err
	}

	result = make(app.TodoCollectionMap)

	for _, row := range rows {
		item, exists := result[row.StateName]
		if exists && (row.TodoID.Valid && row.TodoName.Valid && row.TodoDescription.Valid &&
			row.TodoItemOrder.Valid) {
			item.Todos = append(item.Todos, app.Todo{
				ID:          int(row.TodoID.Int32),
				ItemOrder:   int(row.TodoItemOrder.Int32),
				Name:        row.TodoName.String,
				Description: row.TodoDescription.String,
			})
			result[row.StateName] = item
		} else {
			var todos []app.Todo
			if row.TodoID.Valid && row.TodoName.Valid && row.TodoDescription.Valid &&
				row.TodoItemOrder.Valid {
				t := app.Todo{
					ID:          int(row.TodoID.Int32),
					ItemOrder:   int(row.TodoItemOrder.Int32),
					Name:        row.TodoName.String,
					Description: row.TodoDescription.String,
				}

				todos = append(todos, t)
			}
			result[row.StateName] = app.ToDoCollection{
				TodoState: app.TodoState{
					ID:        row.StateID,
					Name:      row.StateName,
					ItemOrder: row.StateItemOrder,
				},
				Todos: todos,
			}
		}
	}

	return result, nil
}
