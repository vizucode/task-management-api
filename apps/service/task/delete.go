package task

import (
	"context"
	"net/http"

	"task-management-api/apps/middlewares/security"
	"task-management-api/helpers/constants/httpstd"

	"github.com/google/uuid"
	"github.com/vizucode/gokit/logger"
	"github.com/vizucode/gokit/utils/errorkit"
)

func (ts *taskService) DeleteTask(ctx context.Context, taskId string) (err error) {
	// Extract user context
	userContext, ok := security.ExtractUserContext(ctx)
	if !ok {
		err = errorkit.NewErrorStd(http.StatusUnauthorized, "", "User context not found")
		logger.Log.Error(ctx, err)
		return err
	}

	// Validate taskId
	if taskId == "" {
		err = errorkit.NewErrorStd(http.StatusBadRequest, "", "Task ID is required")
		logger.Log.Error(ctx, err)
		return err
	}

	// Validate taskId format
	_, err = uuid.Parse(taskId)
	if err != nil {
		err = errorkit.NewErrorStd(http.StatusBadRequest, "", "Invalid task ID format")
		logger.Log.Error(ctx, err)
		return err
	}

	// Get task detail to verify ownership
	task, err := ts.db.GetDetailTask(ctx, taskId)
	if err != nil {
		logger.Log.Error(ctx, err)
		return err
	}

	// Check if task exists
	if task.Id == uuid.Nil {
		err = errorkit.NewErrorStd(http.StatusNotFound, "", "Task not found")
		logger.Log.Error(ctx, err)
		return err
	}

	// Verify task ownership - pastikan task milik user yang sedang login
	if task.UserId.String() != userContext.Id {
		err = errorkit.NewErrorStd(http.StatusForbidden, "", "You don't have permission to delete this task")
		logger.Log.Error(ctx, err)
		return err
	}

	// Delete task from database
	err = ts.db.DeleteTask(ctx, taskId)
	if err != nil {
		logger.Log.Error(ctx, err)
		return errorkit.NewErrorStd(http.StatusInternalServerError, "", httpstd.InternalServerError)
	}

	return nil
}
