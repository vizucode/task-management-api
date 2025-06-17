package models

import "github.com/google/uuid"

type Todo struct {
	Id          uuid.UUID `gorm:"column:id;primaryKey;autoIncrement"`
	Title       string    `gorm:"column:title"`
	Description string    `gorm:"column:description"`
}
