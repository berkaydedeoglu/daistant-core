package model

import (
	"time"

	"gorm.io/gorm"
)

type ThirdPartyConnection struct {
	gorm.Model
	ID uint

	UserID uint

	Provider     string
	AccessToken  string
	RefreshToken string
	ExpiresAt    time.Time
	Scope        string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
}
