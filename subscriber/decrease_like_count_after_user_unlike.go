package subscriber

import (
	"context"
	"go-food-delivery/common"
	"go-food-delivery/component"
	"go-food-delivery/modules/restaurant/restaurantstorage"
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
