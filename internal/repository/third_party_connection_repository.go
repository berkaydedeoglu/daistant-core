package repository

import (
	"daistant-core/internal/model"
)

func (r *repository) CreateThirdPartyConnection(
	thirdPartyConnection *model.ThirdPartyConnection,
) error {
	return r.db.Create(thirdPartyConnection).Error
}

func (r *repository) GetThirdPartyConnectionByID(
	id uint,
) (*model.ThirdPartyConnection, error) {
	var thirdPartyConnection model.ThirdPartyConnection
	if err := r.db.First(&thirdPartyConnection, id).Error; err != nil {
		return nil, err
	}
	return &thirdPartyConnection, nil
}
