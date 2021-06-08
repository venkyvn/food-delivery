package restaurantlikestorage

import (
	"context"
	"go-food-delivery/common"
	"go-food-delivery/modules/restaurantlike/model"
)

func (s sqlStore) Create(ctx context.Context, data *restaurantlikemodel.RestaurantLike) error {

	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
