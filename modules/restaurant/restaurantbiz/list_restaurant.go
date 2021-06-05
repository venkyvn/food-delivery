package restaurantbiz

import (
	"context"
	"go-food-delivery/common"
	"go-food-delivery/modules/restaurant/restaurantmodel"
	"log"
)

type ListRestaurantStorage interface {
	ListDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		filter *restaurantmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]restaurantmodel.Restaurant, error)
}

type RestaurantLikeStorage interface {
	FetchRestaurantLike(
		ctx context.Context,
		ids []int,
	) (map[int]int, error)
}

type listRestaurantBiz struct {
	store               ListRestaurantStorage
	restaurantLikeStore RestaurantLikeStorage
}

func NewListRestaurantBiz(store ListRestaurantStorage, restaurantLikeStore RestaurantLikeStorage) *listRestaurantBiz {
	return &listRestaurantBiz{
		store:               store,
		restaurantLikeStore: restaurantLikeStore,
	}
}

func (biz *listRestaurantBiz) ListDataByCondition(
	ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
) ([]restaurantmodel.Restaurant, error) {

	result, err := biz.store.ListDataByCondition(ctx, nil, filter, paging, "User")

	if err != nil {
		return nil, common.ErrCannotGetEntity(restaurantmodel.EntityName, err)
	}

	ids := make([]int, len(result))

	for i := range result {
		ids[i] = result[i].Id
	}

	resLikeMap, err := biz.restaurantLikeStore.FetchRestaurantLike(ctx, ids)

	if err != nil {
		log.Printf("cannot fetch restaurant like")
	}

	if resLikeMap != nil {
		for i := range result {
			result[i].LikedCount = resLikeMap[result[i].Id]
		}
	}

	return result, nil
}
