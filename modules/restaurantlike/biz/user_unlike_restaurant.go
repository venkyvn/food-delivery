package restaurantlikedbiz

import (
	"context"
	"go-food-delivery/common"
	"go-food-delivery/modules/restaurantlike/model"
)

type UnlikeRestaurantStorage interface {
	Find(ctx context.Context, data *restaurantlikemodel.RestaurantLike) (*restaurantlikemodel.RestaurantLike, error)
	Delete(ctx context.Context, data *restaurantlikemodel.RestaurantLike) error
}

type unlikeRestaurantBiz struct {
	store UnlikeRestaurantStorage
}

func NewUnlikeRestaurantBiz(store UnlikeRestaurantStorage) *unlikeRestaurantBiz {
	return &unlikeRestaurantBiz{
		store: store,
	}
}

func (biz *unlikeRestaurantBiz) UnlikeRestaurant(ctx context.Context, data *restaurantlikemodel.RestaurantLike) error {

	_, err := biz.store.Find(ctx, data)

	if err != nil {
		if err == common.RecordNotFound {
			return restaurantlikemodel.ErrUserDoesntLikeThisRestaurantYet(err)
		}

		return restaurantlikemodel.ErrUserCannotUnLikeThisRestaurant(err)
	}

	if err := biz.store.Delete(ctx, data); err != nil {
		return restaurantlikemodel.ErrUserCannotUnLikeThisRestaurant(err)
	}

	return nil
}
