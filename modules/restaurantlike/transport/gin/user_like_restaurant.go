package ginrestaurantliked

import (
	"github.com/gin-gonic/gin"
	"go-food-delivery/common"
	"go-food-delivery/component"
	"go-food-delivery/modules/restaurantlike/biz"
	"go-food-delivery/modules/restaurantlike/model"
	"go-food-delivery/modules/restaurantlike/storage"
	"net/http"
)

func UserLikeRestaurant(ctx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		restaurantUid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		data := restaurantlikemodel.RestaurantLike{
			RestaurantId: int(restaurantUid.GetLocalID()),
			UserId:       c.MustGet(common.CurrentUser).(common.Requester).GetUserId(),
		}

		store := restaurantlikestorage.NewSqlStore(ctx.GetMainDBConnection())
		//likeStore := restaurantstorage.NewSQLStore(ctx.GetMainDBConnection())
		biz := restaurantlikedbiz.NewUserLikeRestaurantBiz(store, ctx.GetPubSub())

		if err := biz.LikeRestaurant(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
