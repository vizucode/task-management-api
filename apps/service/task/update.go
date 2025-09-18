package task

import (
	"context"
	"net/http"
	"time"

	"task-management-api/apps/domain"
	"task-management-api/apps/models"
	contextkeys "task-management-api/helpers/constants/context_keys"
	"task-management-api/helpers/constants/httpstd"

	"github.com/google/uuid"
	"github.com/vizucode/gokit/logger"
	"github.com/vizucode/gokit/utils/errorkit"
)

func (ts *taskService) UpdateTask(ctx context.Context, payload []domain.Task) (err error) {
	userContext, ok := ctx.Value(contextkeys.UserContext).(domain.UserContext)
	if !ok {
		err = errorkit.NewErrorStd(http.StatusUnauthorized, "", "User context not found")
		logger.Log.Error(ctx, err)
		return err
	}

	// Validate payload is not empty
	if len(payload) == 0 {
		err = errorkit.NewErrorStd(http.StatusBadRequest, "", "No tasks provided for update")
		logger.Log.Error(ctx, err)
		return err
	}

	// Parse user ID from context
	userID, err := uuid.Parse(userContext.Id)
	if err != nil {
		err = errorkit.NewErrorStd(http.StatusBadRequest, "", "Invalid user ID")
		logger.Log.Error(ctx, err)
		return err
	}

	// Process each task in the payload
	for i, task := range payload {
		// Validate task ID format
		taskUUID, err := uuid.Parse(task.ID)
		if err != nil {
			err = errorkit.NewErrorStd(http.StatusBadRequest, "", "Invalid task ID format for task at index "+string(rune(i)))
			logger.Log.Error(ctx, err)
			return err
		}

		// Get existing task to verify ownership
		existingTask, err := ts.db.GetDetailTask(ctx, task.ID)
		if err != nil {
			logger.Log.Error(ctx, err)
			return err
		}

		// Check if task exists
		if existingTask.Id == uuid.Nil {
			err = errorkit.NewErrorStd(http.StatusNotFound, "", "Task not found for ID: "+task.ID)
			logger.Log.Error(ctx, err)
			return err
		}

		// Verify task ownership - pastikan setiap task milik user yang sedang login
		if existingTask.UserId.String() != userContext.Id {
			err = errorkit.NewErrorStd(http.StatusForbidden, "", "You don't have permission to update task with ID: "+task.ID)
			logger.Log.Error(ctx, err)
			return err
		}

		// Convert domain.Task to models.Task for update
		taskModel := models.Task{
			Id:          taskUUID,
			UserId:      userID, // Pastikan UserId tetap sama dengan user yang login
			Title:       task.Title,
			Description: task.Description,
			Status:      task.Status,
			CreatedAt:   existingTask.CreatedAt, // Preserve original creation time
			UpdatedAt:   time.Now(),             // Update timestamp
		}

		// Update task in database
		err = ts.db.UpdateTask(ctx, taskModel)
		if err != nil {
			logger.Log.Error(ctx, err)
			return errorkit.NewErrorStd(http.StatusInternalServerError, "", httpstd.InternalServerError)
		}
	}

	return nil
}
