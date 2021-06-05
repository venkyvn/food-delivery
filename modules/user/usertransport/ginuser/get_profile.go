package ginuser

import (
	"github.com/gin-gonic/gin"
	"go-food-delivery/common"
	"go-food-delivery/component"
	"net/http"
)

func GetProfile(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet(common.CurrentUser).(common.Requester)

		c.JSON(http.StatusOK, user)
	}
}
