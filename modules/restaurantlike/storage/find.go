package restaurantlikestorage

import (
	"context"
	"go-food-delivery/common"
	"go-food-delivery/modules/restaurantlike/model"
	"gorm.io/gorm"
)

func (s *sqlStore) Find(
	ctx context.Context,
	data *restaurantlikemodel.RestaurantLike,
) (*restaurantlikemodel.RestaurantLike, error) {

	db := s.db

	var result restaurantlikemodel.RestaurantLike

	if err := db.Where("restaurant_id = ? and user_id = ?", data.RestaurantId, data.UserId).
		First(&result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}

		return nil, common.ErrDB(err)
	}

	return &result, nil
}
