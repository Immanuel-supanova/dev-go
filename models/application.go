package models

import (
	"time"

	"github.com/google/uuid"
)

type Application struct {
	ID        uint
	UUID      uuid.UUID `gorm:"unique"`
	Name      string
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DevID     uuid.UUID
}
