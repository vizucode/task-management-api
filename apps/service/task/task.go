package task

import (
	"task-management-api/apps/repositories"

	"github.com/go-playground/validator/v10"
)

type taskService struct {
	db        repositories.IDatabase
	validator validator.Validate
}

func NewTask(
	db repositories.IDatabase,
	validator *validator.Validate,
) *taskService {
	return &taskService{
		db:        db,
		validator: *validator,
	}
}
