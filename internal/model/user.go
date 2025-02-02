package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID    uint
	Name  string `gorm:"not null"`
	Email string `gorm:"not null"`

	ThirdPartyConnections []ThirdPartyConnection

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
