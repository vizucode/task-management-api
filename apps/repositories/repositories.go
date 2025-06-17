package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/origamilabsid/backend-boilerplate/apps/models"
)

type IDatabase interface {
	GetTodo(ctx context.Context, id uuid.UUID) (resp models.Todo, err error)
}
