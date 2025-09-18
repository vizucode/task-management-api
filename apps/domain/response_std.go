package domain

// Metadata represents pagination metadata
type Metadata struct {
	TotalData   int `json:"total_data" example:"100"`
	TotalPage   int `json:"total_page" example:"10"`
	CurrentPage int `json:"current_page" example:"1"`
}

// ResponseJson represents standard API response format
type ResponseJson struct {
	StatusCode string      `json:"status_code" example:"200"`
	Data       interface{} `json:"data"`
	Message    string      `json:"message" example:"Success"`
	Metadata   Metadata    `json:"metadata"`
}
