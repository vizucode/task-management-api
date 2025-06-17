package todo

import (
	"github.com/go-playground/validator/v10"
	"github.com/origamilabsid/backend-boilerplate/apps/repositories"
)

type todoService struct {
	db        repositories.IDatabase
	validator validator.Validate
}

func NewTodoService(
	db repositories.IDatabase,
	validator *validator.Validate,
) *todoService {
	return &todoService{
		db:        db,
		validator: *validator,
	}
}
