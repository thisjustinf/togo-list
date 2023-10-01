package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID uuid.UUID `gorm:"not null;"`
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Username  string `gorm:"not null;unique_idx:idx_name"`
	Password  string `gorm:"not null"`
}
