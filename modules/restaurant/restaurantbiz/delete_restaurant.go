package restaurantbiz

import "context"

type DeleteRestaurantStore interface {
	FindDateByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) error

	SoftDeleteData(
		ctx context.Context,
		id int,
	) error
}

type deleteRestaurantBiz struct {
	store DeleteRestaurantStore
}

func NewDeleteRestaurantBiz(store DeleteRestaurantStore) *deleteRestaurantBiz {
	return &deleteRestaurantBiz{
		store: store,
	}
}
