package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title       string    `gorm:"not null;"`
	Description string    `gorm:"not null;"`
	Completed   bool      `gorm:"not null;"`
	UserID      uuid.UUID `gorm:"not null;"`
}
