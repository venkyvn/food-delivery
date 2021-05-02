package restaurantstorage

import (
	"context"
	"go-food-delivery/modules/restaurant/restaurantmodel"
)

func (s *sqlStore) SoftDeleteData(
	ctx context.Context,
	id int,
) error {

	db := s.db
	if err := db.Table(restaurantmodel.Restaurant{}.TableName()).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"status": 0,
		}).
		Error; err != nil {
		return err
	}

	return nil
}
