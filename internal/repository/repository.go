package repository

import (
	"context"
	"daistant-core/internal/model"

	"gorm.io/gorm"
)

type Repository interface {
	GetUserByID(id uint) (*model.User, error)

	CreateThirdPartyConnection(ctx context.Context, thirdPartyConnection *model.ThirdPartyConnection) error
	GetThirdPartyConnectionByID(ctx context.Context, id uint) (*model.ThirdPartyConnection, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}
