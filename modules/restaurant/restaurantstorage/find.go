package restaurantstorage

import (
	"context"
	"go-food-delivery/common"
	"go-food-delivery/modules/restaurant/restaurantmodel"
)

func (s *sqlStore) FindDataByCollection(
	ctx context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*restaurantmodel.Restaurant, error) {

	var result restaurantmodel.Restaurant

	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.Where(condition).First(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return &result, nil
}
