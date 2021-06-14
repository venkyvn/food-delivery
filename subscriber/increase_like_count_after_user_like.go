package subscriber

import (
	"context"
	"go-food-delivery/common"
	"go-food-delivery/component"
	"go-food-delivery/modules/restaurant/restaurantstorage"
	"go-food-delivery/pubsub"
	"go-food-delivery/skio"
)

type HasRestaurantId interface {
	GetRestaurantId() int
	GetUserId() int
}

func IncreaseLikeCountAfterUserLikeRestaurant(appContext component.AppContext, context context.Context) {
	c, _ := appContext.GetPubSub().Subscribe(context, common.TopicUserLikeRestaurant)

	store := restaurantstorage.NewSQLStore(appContext.GetMainDBConnection())

	go func() {
		for {
			msg := <-c
			likeData := msg.Data().(HasRestaurantId)
			_ = store.IncreaseLikeCount(context, likeData.GetRestaurantId())
		}
	}()
}

// I wish i could do something like that
//func RunIncreaseLikeAfterUserLikeRestaurant(appContext component.AppContext) func(ctx context.Context, message *pubsub.Message) error {
//	store := restaurantstorage.NewSQLStore(appContext.GetMainDBConnection())
//
//	return func(ctx context.Context, message *pubsub.Message) error {
//		likeData := message.Data().(HasRestaurantId)
//		return store.IncreaseLikeCount(ctx, likeData.GetRestaurantId())
//	}
//}

func RunIncreaseLikeCountAfterUserLikeRestaurant(appContext component.AppContext) consumerJob {
	return consumerJob{
		Title: "Increase Like After User Like Restaurant",
		Handler: func(ctx context.Context, message *pubsub.Message) error {
			store := restaurantstorage.NewSQLStore(appContext.GetMainDBConnection())
			likeData := message.Data().(HasRestaurantId)
			return store.IncreaseLikeCount(ctx, likeData.GetRestaurantId())
		},
	}
}

func EmitRealTimeAfterUserLikeRestaurant(rtEngine skio.RealtimeEngine) consumerJob {
	return consumerJob{
		Title: "Emit real time after User Like Restaurant",
		Handler: func(ctx context.Context, message *pubsub.Message) error {
			likeData := message.Data().(HasRestaurantId)
			return rtEngine.EmitToUser(likeData.GetUserId(), string(message.Topic()), likeData)
		},
	}
}
