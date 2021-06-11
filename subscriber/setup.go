package subscriber

import (
	"context"
	"go-food-delivery/component"
)

// not using anymore. but I still put it there to easily follow it
func Setup(appCtx component.AppContext) {
	IncreaseLikeCountAfterUserLikeRestaurant(appCtx, context.Background())
	DecreaseLikeCountAfterUserUnlike(appCtx, context.Background())
}
