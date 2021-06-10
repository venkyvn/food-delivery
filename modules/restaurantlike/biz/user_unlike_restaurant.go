package restaurantlikedbiz

import (
	"context"
	"go-food-delivery/common"
	"go-food-delivery/modules/restaurantlike/model"
	"go-food-delivery/pubsub"
)

type UnlikeRestaurantStorage interface {
	Find(ctx context.Context, data *restaurantlikemodel.RestaurantLike) (*restaurantlikemodel.RestaurantLike, error)
	Delete(ctx context.Context, data *restaurantlikemodel.RestaurantLike) error
}

//type DecreaseLikeCountStorage interface {
//	DecreaseLikeCount(ctx context.Context, restaurantId int) error
//}

type unlikeRestaurantBiz struct {
	store UnlikeRestaurantStorage
	//decreaseStore DecreaseLikeCountStorage
	pubSubStore pubsub.PubSub
}

func NewUnlikeRestaurantBiz(
	store UnlikeRestaurantStorage,
	//decreaseStore DecreaseLikeCountStorage,
	pubSubStore pubsub.PubSub,
) *unlikeRestaurantBiz {
	return &unlikeRestaurantBiz{
		store: store,
		//decreaseStore: decreaseStore,
		pubSubStore: pubSubStore,
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

	// Side effect
	biz.pubSubStore.Publish(ctx, common.TopicUserUnLikeRestaurant, pubsub.NewMessage(data))
	//if err := biz.decreaseStore.DecreaseLikeCount(ctx, data.RestaurantId); err != nil {
	//	log.Println("cannot decrease like count of restaurant", err)
	//}

	return nil
}
