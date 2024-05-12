package models

import (
	"time"

	"github.com/google/uuid"
)

type Developer struct {
	ID           uint
	UUID         uuid.UUID `gorm:"unique"`
	Email        string    `gorm:"unique"`
	Password     string
	IsActive     bool
	IsAdmin      bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Applications []Application `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:DevID"`
}
