package restaurantlikestorage

import (
	"context"
	"go-food-delivery/common"
	restaurantlikemodel "go-food-delivery/modules/restaurantlike/model"
)

func (s *sqlStore) FetchRestaurantLike(ctx context.Context, ids []int) (map[int]int, error) {

	result := make(map[int]int)

	type sqlData struct {
		RestaurantId int `gorm:"column:restaurant_id;"`
		Count        int `gorm:"column:count;"`
	}

	var likeResult []sqlData

	if err := s.db.
		Table(restaurantlikemodel.RestaurantLike{}.TableName()).
		Select("restaurant_id, count(restaurant_id) as count").
		Where("restaurant_id in (?)", ids).
		Group("restaurant_id").
		Find(&likeResult).
		Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for _, item := range likeResult {
		result[item.RestaurantId] = item.Count
	}

	return result, nil
}
