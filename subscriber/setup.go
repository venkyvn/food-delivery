package subscriber

import (
	"context"
	"go-food-delivery/component"
)

func Setup(appCtx component.AppContext) {
	IncreaseLikeCountAfterUserLikeRestaurant(appCtx, context.Background())
	DecreaseLikeCountAfterUserUnlike(appCtx, context.Background())
}
