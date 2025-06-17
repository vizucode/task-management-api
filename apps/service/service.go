package service

import (
	"context"

	"github.com/origamilabsid/backend-boilerplate/apps/domain"
)

type ITodoService interface {
	GetTodo(ctx context.Context) (resp []domain.Todo, err error)
}
