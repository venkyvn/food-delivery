package restaurantlikedbiz

import (
	"context"
	"go-food-delivery/common"
	"go-food-delivery/modules/restaurantlike/model"
)

type UserLikedRestaurantStorage interface {
	GetUserLikedRestaurantById(
		ctx context.Context,
		conditions map[string]interface{},
		filter *restaurantlikemodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]common.SimpleUser, error)
}

type restaurantLikedBiz struct {
	store UserLikedRestaurantStorage
}

func NewRestaurantLikedBiz(store UserLikedRestaurantStorage) *restaurantLikedBiz {
	return &restaurantLikedBiz{
		store: store,
	}
}

func (biz *restaurantLikedBiz) GetUserLikedRestaurant(ctx context.Context,
	filter *restaurantlikemodel.Filter,
	paging *common.Paging) ([]common.SimpleUser, error) {

	result, err := biz.store.GetUserLikedRestaurantById(ctx, nil, filter, paging)

	if err != nil {
		common.ErrCannotListEntity(restaurantlikemodel.EntityName, err)
	}

	return result, nil
}
