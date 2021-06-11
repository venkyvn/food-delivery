package subscriber

import (
	"context"
	"go-food-delivery/common"
	"go-food-delivery/component"
	"go-food-delivery/modules/restaurant/restaurantstorage"
	"go-food-delivery/pubsub"
)

func DecreaseLikeCountAfterUserUnlike(appCtx component.AppContext, context context.Context) {

	store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
	c, _ := appCtx.GetPubSub().Subscribe(context, common.TopicUserUnLikeRestaurant)

	go func() {
		for {
			msg := <-c
			restaurantId := msg.Data().(HasRestaurantId)

			store.DecreaseLikeCount(context, restaurantId.GetRestaurantId())
		}
	}()
}

func RunDecreaseLikeCountAfterUserLikeRestaurant(appContext component.AppContext) consumerJob {
	return consumerJob{
		Title: "Decrease like count after user unlike restaurant",
		Handler: func(ctx context.Context, message *pubsub.Message) error {
			store := restaurantstorage.NewSQLStore(appContext.GetMainDBConnection())
			restaurantId := message.Data().(HasRestaurantId)

			return store.DecreaseLikeCount(ctx, restaurantId.GetRestaurantId())
		},
	}
}
