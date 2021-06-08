package restaurantlikestorage

import (
	"context"
	"go-food-delivery/common"
	"go-food-delivery/modules/restaurantlike/model"
)

func (s *sqlStore) Delete(ctx context.Context, data *restaurantlikemodel.RestaurantLike) error {

	db := *s.db

	if err := db.Table(data.TableName()).
		Where("restaurant_id = ? and user_id = ?", data.RestaurantId, data.UserId).
		Delete(nil).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
