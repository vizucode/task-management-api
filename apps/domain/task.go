package domain

// Task represents a task entity
type Task struct {
	ID          string `json:"id" example:"123e4567-e89b-12d3-a456-426614174000"`
	UserId      string `json:"user_id"`
	Title       string `json:"title" example:"Complete project documentation" validate:"required"`
	Description string `json:"description" example:"Write comprehensive API documentation"`
	Status      string `json:"status" example:"pending" validate:"required" enums:"pending,in_progress,completed"`
	CreatedAt   string `json:"created_at" example:"2024-01-15 10:30"`
}

type TaskResponse struct {
	ID          string `json:"id" example:"123e4567-e89b-12d3-a456-426614174000"`
	Title       string `json:"title" example:"Complete project documentation" validate:"required"`
	Description string `json:"description" example:"Write comprehensive API documentation"`
	Status      string `json:"status" example:"pending" validate:"required" enums:"pending,in_progress,completed"`
	CreatedAt   string `json:"created_at" example:"2024-01-15 10:30"`
}
