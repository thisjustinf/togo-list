package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Togo struct {
	ID uuid.UUID `gorm:"type:uuid;"`
}

type Togos struct {
}

func (togo *Togo) BeforeCreate(tx *gorm.DB) (err error) {
	togo.ID = uuid.New()
	return
}
