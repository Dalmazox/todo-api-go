package services

import (
	"context"

	"github.com/dalmazox/todo-api/internal/db"
)

type TodoService struct {
	repository *db.Repository
}

func NewTodoService(repository *db.Repository) *TodoService {
	return &TodoService{
		repository: repository,
	}
}

func (t *TodoService) CreateTodo(ctx context.Context, description string, done bool) error {
	return t.repository.CreateTodo(ctx, description, done)
}
