package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"go-food-delivery/component"
)

func DeleteRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		//id, err := strconv.Atoi(c.Param("id"))
		//
		//if err != nil {
		//	panic(common.ErrInvalidRequest(err))
		//}
		//
		//store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		//biz := restaurantbiz.NewDeleteRestaurantBiz(store)
	}
}
