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

func GetUserLikedRestaurant(ctx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		restaurantUID, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		filter := restaurantlikemodel.Filter{
			RestaurantId: int(restaurantUID.GetLocalID()),
		}

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		db := ctx.GetMainDBConnection()
		store := restaurantlikestorage.NewSqlStore(db)
		biz := restaurantlikedbiz.NewRestaurantLikedBiz(store)

		result, err := biz.GetUserLikedRestaurant(c.Request.Context(), &filter, &paging)

		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
