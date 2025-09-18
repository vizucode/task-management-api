package repositories

import (
	"context"

	"task-management-api/apps/models"
)

type IDatabase interface {
	CreateUser(ctx context.Context, user models.User) (err error)
	GetUserByEmail(ctx context.Context, email string) (resp models.User, err error)
	GetTaskList(ctx context.Context, filter models.Filter) (resp []models.Task, err error)
	CountTaskList(ctx context.Context, filter models.Filter) (resp int64, err error)
	GetDetailTask(ctx context.Context, taskId string) (resp models.Task, err error)
	CreateTask(ctx context.Context, payload models.Task) (err error)
	UpdateTask(ctx context.Context, payload models.Task) (err error)
	DeleteTask(ctx context.Context, taskId string) (err error)
}
