package task

import (
	"context"
	"net/http"
	"time"

	"task-management-api/apps/domain"
	"task-management-api/apps/models"
	"task-management-api/helpers/constants/httpstd"

	"github.com/google/uuid"
	"github.com/vizucode/gokit/logger"
	"github.com/vizucode/gokit/utils/errorkit"
)

func (ts *taskService) CreateTask(ctx context.Context, payload domain.Task) (err error) {
	// Validate required fields
	if payload.Title == "" {
		err = errorkit.NewErrorStd(http.StatusBadRequest, "", "Title is required")
		logger.Log.Error(ctx, err)
		return err
	}

	if payload.Status == "" {
		err = errorkit.NewErrorStd(http.StatusBadRequest, "", "Status is required")
		logger.Log.Error(ctx, err)
		return err
	}

	userID, err := uuid.Parse(payload.UserId)
	if err != nil {
		err = errorkit.NewErrorStd(http.StatusBadRequest, "", "Invalid user ID")
		logger.Log.Error(ctx, err)
		return err
	}

	taskModel := models.Task{
		Id:          uuid.New(), // Generate new UUID for task
		UserId:      userID,
		Title:       payload.Title,
		Description: payload.Description,
		Status:      payload.Status,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// Create task in database
	err = ts.db.CreateTask(ctx, taskModel)
	if err != nil {
		logger.Log.Error(ctx, err)
		return errorkit.NewErrorStd(http.StatusInternalServerError, "", httpstd.InternalServerError)
	}

	return nil
}
