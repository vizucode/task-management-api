package domain

type ResponseJson struct {
	StatusCode string      `json:"status_code"`
	Data       interface{} `json:"data"`
	Message    string      `json:"message"`
}
