package models

type Filter struct {
	Page       int
	Limit      int
	TaskStatus string
	UserId     string
}
