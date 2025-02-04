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
	if err := r.db.WithContext(ctx).First(&thirdPartyConnection, id).Error; err != nil {
		return nil, err
	}
	return &thirdPartyConnection, nil
}

func (r *repository) GetThirdPartyConnectionByUserID(
	ctx context.Context,
	userID uint,
	thirdPartyType string,
) (*model.ThirdPartyConnection, error) {
	var tpc model.ThirdPartyConnection
	err := r.db.WithContext(ctx).
		Where("user_id = ? AND provider = ?", userID, thirdPartyType).
		First(&tpc).Error
	if err != nil {
		return nil, err
	}
	return &tpc, nil
}

func (r *repository) UpdateThirdPartyConnection(ctx context.Context, thirdPartyConnection *model.ThirdPartyConnection) error {
	return r.db.
		WithContext(ctx).
		Model(&model.ThirdPartyConnection{}).
		Where("user_id = ? AND provider = ?", thirdPartyConnection.UserID, thirdPartyConnection.Provider).
		Updates(thirdPartyConnection).Error
}
