package rest

import (
	"net/http"
	"strconv"

	"task-management-api/apps/domain"
	"task-management-api/apps/middlewares/security"

	"github.com/gofiber/fiber/v2"
	"github.com/vizucode/gokit/logger"
)

// GetTaskList godoc
// @Summary Get list of tasks
// @Description Get paginated list of tasks for authenticated user
// @Tags Tasks
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Param status query string false "Filter by task status"
// @Success 200 {object} domain.ResponseJson{data=[]domain.Task} "Successfully get task list"
// @Failure 401 {object} domain.ResponseJson "Unauthorized"
// @Failure 500 {object} domain.ResponseJson "Internal server error"
// @Router /tasks [get]
func (r *rest) GetTaskList(c *fiber.Ctx) error {
	// Extract user context
	ctx, userContext, ok := security.ExtractUserContextFiber(c)
	if !ok {
		return r.ResponseJson(c, http.StatusUnauthorized, nil, domain.Metadata{}, "User context not found")
	}

	// Parse query parameters
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	status := c.Query("status", "")

	filter := domain.Filter{
		Page:       page,
		Limit:      limit,
		TaskStatus: status,
		UserId:     userContext.Id,
	}

	// Get task list
	tasks, err := r.task.GetTaskList(ctx, filter)
	if err != nil {
		logger.Log.Error(ctx, err)
		return r.ResponseJson(c, http.StatusInternalServerError, nil, domain.Metadata{}, err.Error())
	}

	// Get total count
	totalCount, err := r.task.CountTaskList(c.Context(), userContext.Id, status)
	if err != nil {
		logger.Log.Error(c.Context(), err)
		return r.ResponseJson(c, http.StatusInternalServerError, nil, domain.Metadata{}, err.Error())
	}

	totalPage := int(totalCount) / limit
	if int(totalCount)%limit > 0 {
		totalPage++
	}

	metadata := domain.Metadata{
		TotalData:   int(totalCount),
		TotalPage:   totalPage,
		CurrentPage: page,
	}

	return r.ResponseJson(c, http.StatusOK, tasks, metadata, "Successfully get task list")
}

// GetTaskDetail godoc
// @Summary Get task detail
// @Description Get detailed information of a specific task
// @Tags Tasks
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Task ID"
// @Success 200 {object} domain.ResponseJson{data=domain.Task} "Successfully get task detail"
// @Failure 401 {object} domain.ResponseJson "Unauthorized"
// @Failure 404 {object} domain.ResponseJson "Task not found"
// Router /tasks/{id} [get]
func (r *rest) GetTaskDetail(c *fiber.Ctx) error {
	// Extract user context
	ctx, _, ok := security.ExtractUserContextFiber(c)
	if !ok {
		return r.ResponseJson(c, http.StatusUnauthorized, nil, domain.Metadata{}, "User context not found")
	}

	taskId := c.Params("id")

	task, err := r.task.TaskDetail(ctx, taskId)
	if err != nil {
		logger.Log.Error(ctx, err)
		return r.ResponseJson(c, http.StatusNotFound, nil, domain.Metadata{}, err.Error())
	}

	return r.ResponseJson(c, http.StatusOK, task, domain.Metadata{}, "Successfully get task detail")
}

// CreateTask godoc
// @Summary Create a new task
// @Description Create a new task for authenticated user
// @Tags Tasks
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param task body domain.Task true "Task data"
// @Success 201 {object} domain.ResponseJson "Task created successfully"
// @Failure 400 {object} domain.ResponseJson "Bad request"
// @Failure 401 {object} domain.ResponseJson "Unauthorized"
// @Router /tasks [post]
func (r *rest) CreateTask(c *fiber.Ctx) error {
	var payload domain.Task

	// Extract user context
	ctx, userContext, ok := security.ExtractUserContextFiber(c)
	if !ok {
		return r.ResponseJson(c, http.StatusUnauthorized, nil, domain.Metadata{}, "User context not found")
	}

	payload.UserId = userContext.Id

	if err := c.BodyParser(&payload); err != nil {
		logger.Log.Error(c.Context(), err)
		return r.ResponseJson(c, http.StatusBadRequest, nil, domain.Metadata{}, "Invalid request body")
	}

	err := r.task.CreateTask(ctx, payload)
	if err != nil {
		logger.Log.Error(c.Context(), err)
		return r.ResponseJson(c, http.StatusBadRequest, nil, domain.Metadata{}, err.Error())
	}

	return r.ResponseJson(c, http.StatusCreated, nil, domain.Metadata{}, "Task created successfully")
}

// UpdateTask godoc
// @Summary Update multiple tasks
// @Description Update multiple tasks for authenticated user
// @Tags Tasks
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param tasks body []domain.Task true "Array of tasks to update"
// @Success 200 {object} domain.ResponseJson "Tasks updated successfully"
// @Failure 400 {object} domain.ResponseJson "Bad request"
// @Failure 401 {object} domain.ResponseJson "Unauthorized"
// @Failure 403 {object} domain.ResponseJson "Forbidden - task ownership violation"
// @Router /tasks [put]
func (r *rest) UpdateTask(c *fiber.Ctx) error {
	var payload []domain.Task

	// Extract user context
	ctx, _, ok := security.ExtractUserContextFiber(c)
	if !ok {
		return r.ResponseJson(c, http.StatusUnauthorized, nil, domain.Metadata{}, "User context not found")
	}

	if err := c.BodyParser(&payload); err != nil {
		logger.Log.Error(ctx, err)
		return r.ResponseJson(c, http.StatusBadRequest, nil, domain.Metadata{}, "Invalid request body")
	}

	err := r.task.UpdateTask(ctx, payload)
	if err != nil {
		logger.Log.Error(ctx, err)
		return r.ResponseJson(c, http.StatusBadRequest, nil, domain.Metadata{}, err.Error())
	}

	return r.ResponseJson(c, http.StatusOK, nil, domain.Metadata{}, "Tasks updated successfully")
}

// DeleteTask godoc
// @Summary Delete a task
// @Description Delete a specific task by ID
// @Tags Tasks
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Task ID"
// @Success 200 {object} domain.ResponseJson "Task deleted successfully"
// @Failure 400 {object} domain.ResponseJson "Bad request"
// @Failure 401 {object} domain.ResponseJson "Unauthorized"
// @Failure 403 {object} domain.ResponseJson "Forbidden - task ownership violation"
// @Failure 404 {object} domain.ResponseJson "Task not found"
// @Router /tasks/{id} [delete]
func (r *rest) DeleteTask(c *fiber.Ctx) error {
	// Extract user context
	ctx, _, ok := security.ExtractUserContextFiber(c)
	if !ok {
		return r.ResponseJson(c, http.StatusUnauthorized, nil, domain.Metadata{}, "User context not found")
	}

	taskId := c.Params("id")

	err := r.task.DeleteTask(ctx, taskId)
	if err != nil {
		logger.Log.Error(ctx, err)
		return r.ResponseJson(c, http.StatusBadRequest, nil, domain.Metadata{}, err.Error())
	}

	return r.ResponseJson(c, http.StatusOK, nil, domain.Metadata{}, "Task deleted successfully")
}
