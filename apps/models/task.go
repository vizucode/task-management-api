package models

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	Id          uuid.UUID `gorm:"column:id;primaryKey;autoIncrement"`
	UserId      uuid.UUID `gorm:"column:id"`
	Title       string    `gorm:"column:title"`
	Description string    `gorm:"column:description"`
	Status      string    `gorm:"column:status"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`

	User User `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
