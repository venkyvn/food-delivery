package restaurantlikedbiz

import (
	"context"
	"go-food-delivery/common"
	"go-food-delivery/modules/restaurantlike/model"
	"go-food-delivery/pubsub"
)

type LikeRestaurantStorage interface {
	Find(ctx context.Context, data *restaurantlikemodel.RestaurantLike) (*restaurantlikemodel.RestaurantLike, error)
	Create(ctx context.Context, data *restaurantlikemodel.RestaurantLike) error
}

//type IncreaseLikeCountStorage interface {
//	IncreaseLikeCount(ctx context.Context, restaurantId int) error
//}

type likeRestaurantBiz struct {
	store LikeRestaurantStorage
	//increaseStore IncreaseLikeCountStorage
	pubSub pubsub.PubSub
}

func NewUserLikeRestaurantBiz(
	store LikeRestaurantStorage,
	//increaseStore IncreaseLikeCountStorage,
	pubSub pubsub.PubSub,
) *likeRestaurantBiz {
	return &likeRestaurantBiz{
		store: store,
		//increaseStore: increaseStore,
		pubSub: pubSub,
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

	//side effect
	biz.pubSub.Publish(ctx, common.TopicUserLikeRestaurant, pubsub.NewMessage(data))

	//if err := biz.increaseStore.IncreaseLikeCount(ctx, data.RestaurantId); err != nil {
	//	log.Println("cannot increase like count ", err)
	//}

	//job := asyncjob.NewJob(func(ctx context.Context) error {
	//	return biz.increaseStore.IncreaseLikeCount(ctx, data.RestaurantId)
	//})

	//_ = asyncjob.NewGroup(true, job).Run(ctx)

	//side effect increase like count
	//go func() {
	//	defer common.AppRecover()
	//	if err := biz.increaseStore.IncreaseLikeCount(ctx, data.RestaurantId); err != nil {
	//		log.Println("cannot increase like count ", err)
	//	}
	//}()

	return nil
}
