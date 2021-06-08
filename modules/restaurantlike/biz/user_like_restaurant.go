package restaurantlikedbiz

import (
	"context"
	"go-food-delivery/common"
	"go-food-delivery/modules/restaurantlike/model"
)

type LikeRestaurantStorage interface {
	Find(ctx context.Context, data *restaurantlikemodel.RestaurantLike) (*restaurantlikemodel.RestaurantLike, error)
	Create(ctx context.Context, data *restaurantlikemodel.RestaurantLike) error
}

type likeRestaurantBiz struct {
	store LikeRestaurantStorage
}

func NewUserLikeRestaurantBiz(store LikeRestaurantStorage) *likeRestaurantBiz {
	return &likeRestaurantBiz{
		store: store,
	}
}

func (biz *likeRestaurantBiz) LikeRestaurant(ctx context.Context, data *restaurantlikemodel.RestaurantLike) error {
	userLiked, err := biz.store.Find(ctx, data)

	if err != nil && err != common.RecordNotFound {
		return restaurantlikemodel.ErrUserCannotLikeThisRestaurant(err)
	}

	if userLiked != nil {
		return restaurantlikemodel.ErrUserAlreadyLikeThisRestaurant(err)
	}

	if err := biz.store.Create(ctx, data); err != nil {
		return restaurantlikemodel.ErrUserCannotLikeThisRestaurant(err)
	}

	return nil
}
