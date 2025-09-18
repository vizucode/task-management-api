package domain

// Filter represents filtering and pagination parameters
type Filter struct {
	Page       int    `json:"page" example:"1"`
	Limit      int    `json:"limit" example:"10"`
	TaskStatus string `json:"task_status" example:"pending"`
	UserId     string `json:"user_id" example:"123e4567-e89b-12d3-a456-426614174000"`
}
