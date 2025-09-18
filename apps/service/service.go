package service

import (
	"context"
	"task-management-api/apps/domain"
)

type ITaskService interface {
	GetTaskList(ctx context.Context, filter domain.Filter) (resp []domain.TaskResponse, err error)
	TaskDetail(ctx context.Context, id string) (resp domain.TaskResponse, err error)
	CountTaskList(ctx context.Context, id string, status string) (resp int64, err error)
	CreateTask(ctx context.Context, payload domain.Task) (err error)
	UpdateTask(ctx context.Context, payload []domain.Task) (err error)
	DeleteTask(ctx context.Context, taskId string) (err error)
}

type IUserService interface {
	SignUp(ctx context.Context, payload domain.User) (err error)
	SignIn(ctx context.Context, payload domain.User) (resp domain.SignInResponse, err error)
}
