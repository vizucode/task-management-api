package task

import (
	"context"
	"net/http"

	"task-management-api/apps/domain"
	"task-management-api/apps/models"

	"github.com/google/uuid"
	"github.com/vizucode/gokit/logger"
	"github.com/vizucode/gokit/utils/errorkit"
)

func (ts *taskService) GetTaskList(ctx context.Context, filter domain.Filter) (resp []domain.TaskResponse, err error) {
	modelFilter := models.Filter{
		Page:       filter.Page,
		Limit:      filter.Limit,
		TaskStatus: filter.TaskStatus,
		UserId:     filter.UserId,
	}

	tasks, err := ts.db.GetTaskList(ctx, modelFilter)
	if err != nil {
		logger.Log.Error(ctx, err)
		return resp, err
	}

	resp = make([]domain.TaskResponse, len(tasks))
	for i, task := range tasks {
		resp[i] = domain.TaskResponse{
			ID:          task.Id.String(),
			Title:       task.Title,
			Description: task.Description,
			Status:      task.Status,
			CreatedAt:   task.CreatedAt.Format("2006-01-02 15:04"),
		}
	}

	return resp, nil
}

func (ts *taskService) TaskDetail(ctx context.Context, id string) (resp domain.TaskResponse, err error) {
	if id == "" {
		err = errorkit.NewErrorStd(http.StatusBadRequest, "", "Task ID is required")
		logger.Log.Error(ctx, err)
		return resp, err
	}

	task, err := ts.db.GetDetailTask(ctx, id)
	if err != nil {
		logger.Log.Error(ctx, err)
		return resp, err
	}

	if task.Id == uuid.Nil {
		err = errorkit.NewErrorStd(http.StatusNotFound, "", "Task not found")
		logger.Log.Error(ctx, err)
		return resp, err
	}

	resp = domain.TaskResponse{
		ID:          task.Id.String(),
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		CreatedAt:   task.CreatedAt.Format("2006-01-02 15:04"),
	}

	return resp, nil
}

func (ts *taskService) CountTaskList(ctx context.Context, userId string, status string) (resp int64, err error) {
	if userId == "" {
		err = errorkit.NewErrorStd(http.StatusBadRequest, "", "User ID is required")
		logger.Log.Error(ctx, err)
		return resp, err
	}

	modelFilter := models.Filter{
		TaskStatus: status,
		UserId:     userId,
	}

	count, err := ts.db.CountTaskList(ctx, modelFilter)
	if err != nil {
		logger.Log.Error(ctx, err)
		return resp, err
	}

	return count, nil
}
