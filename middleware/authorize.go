package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-food-delivery/common"
	"go-food-delivery/component"
	"go-food-delivery/component/tokenprovider/jwt"
	"go-food-delivery/modules/user/usermodel"
	"go-food-delivery/modules/user/userstorage"
	"strings"
)

func RequireAuth(appCtx component.AppContext) func(ctx *gin.Context) {
	tokenProvider := jwt.NewJwtProvider(appCtx.GetSecretKey())
	return func(c *gin.Context) {
		token, err := extractTokenFromHeaderString(c.GetHeader("Authorization"))
		if err != nil {
			panic(err)
		}

		payload, err := tokenProvider.Validate(token)

		if err != nil {
			panic(err)
		}

		db := appCtx.GetMainDBConnection()
		store := userstorage.NewSqlStore(db)

		user, err := store.FindUser(c.Request.Context(), map[string]interface{}{"id": payload.UserId})

		if err != nil {
			panic(err)
		}

		if user.Status == 0 {
			panic(common.ErrNoPermission(usermodel.EntityName, err))
		}

		user.Mask(false)

		c.Set(common.CurrentUser, user)
		c.Next()
	}
}

func extractTokenFromHeaderString(header string) (string, error) {
	parts := strings.Split(header, " ")

	if len(parts) < 2 || parts[0] != "Bearer" || strings.TrimSpace(parts[1]) == "" {
		return "", ErrWrongAuthHeader(nil)
	}

	return parts[1], nil
}

func ErrWrongAuthHeader(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("wrong authen header"),
		fmt.Sprintf("ErrWrongAuthHeader"),
	)

}
