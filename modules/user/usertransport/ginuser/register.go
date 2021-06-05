package ginuser

import (
	"github.com/gin-gonic/gin"
	"go-food-delivery/common"
	"go-food-delivery/component"
	"go-food-delivery/component/hasher"
	"go-food-delivery/modules/user/userbiz"
	"go-food-delivery/modules/user/usermodel"
	"go-food-delivery/modules/user/userstorage"
	"net/http"
)

func Register(appCtx component.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		var data usermodel.UserCreate

		if err := c.ShouldBind(&data); err != nil {
		}

		store := userstorage.NewSqlStore(db)
		md5 := hasher.NewMd5Hash()
		biz := userbiz.NewRegistrationBiz(store, md5)

		if err := biz.Register(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
