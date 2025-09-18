package user

import (
	"task-management-api/apps/repositories"

	"github.com/go-playground/validator/v10"
)

type user struct {
	db        repositories.IDatabase
	validator validator.Validate
}

func NewUser(
	db repositories.IDatabase,
	validator *validator.Validate,
) *user {
	return &user{
		db:        db,
		validator: *validator,
	}
}
