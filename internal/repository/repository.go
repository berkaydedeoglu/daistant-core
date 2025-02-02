package repository

import (
	"daistant-core/internal/model"

	"gorm.io/gorm"
)

type Repository interface {
	GetUserByID(id uint) (*model.User, error)

	CreateThirdPartyConnection(thirdPartyConnection *model.ThirdPartyConnection) error
	GetThirdPartyConnectionByID(id uint) (*model.ThirdPartyConnection, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}
