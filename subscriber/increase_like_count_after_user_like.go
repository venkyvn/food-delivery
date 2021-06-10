package subscriber

import (
	"context"
	"go-food-delivery/common"
	"go-food-delivery/component"
	"go-food-delivery/modules/restaurant/restaurantstorage"
)

type HasRestaurantId interface {
	GetRestaurantId() int
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
