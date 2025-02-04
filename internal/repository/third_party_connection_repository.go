package repository

import (
	"context"
	"daistant-core/internal/model"
)

func (r *repository) CreateThirdPartyConnection(ctx context.Context, thirdPartyConnection *model.ThirdPartyConnection) error {
	return r.db.WithContext(ctx).Create(thirdPartyConnection).Error
}

func (r *repository) GetThirdPartyConnectionByID(
	ctx context.Context,
	id uint,
) (*model.ThirdPartyConnection, error) {
	var thirdPartyConnection model.ThirdPartyConnection
	if err := r.db.First(&thirdPartyConnection, id).Error; err != nil {
		return nil, err
	}
	return &thirdPartyConnection, nil
}
