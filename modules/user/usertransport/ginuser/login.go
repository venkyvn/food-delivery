package ginuser

import (
	"github.com/gin-gonic/gin"
	"go-food-delivery/common"
	"go-food-delivery/component"
	"go-food-delivery/component/hasher"
	"go-food-delivery/component/tokenprovider/jwt"
	"go-food-delivery/modules/user/userbiz"
	"go-food-delivery/modules/user/usermodel"
	"go-food-delivery/modules/user/userstorage"
	"net/http"
)

func Login(ctx component.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		var loginUser usermodel.UserLogin

		if err := c.ShouldBind(&loginUser); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := ctx.GetMainDBConnection()

		secretKey := ctx.GetSecretKey()
		tokenProvider := jwt.NewJwtProvider(secretKey)

		md5 := hasher.NewMd5Hash()

		expiry := 60 * 60 * 24 * 7
		store := userstorage.NewSqlStore(db)
		biz := userbiz.NewLoginBiz(store, tokenProvider, md5, expiry)

		token, err := biz.Login(c.Request.Context(), &loginUser)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(token))
	}
}
