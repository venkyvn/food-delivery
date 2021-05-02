package middleware

import (
	"github.com/gin-gonic/gin"
	"go-food-delivery/common"
	"go-food-delivery/component"
)

/*
put app context here
so your app can check which environment is executed
and you can easily handle which error you wanna show
*/
func Recover(ac component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.Header("Content-Type", "application/json")

				// app Err handle in production with app context environment
				if appErr, ok := err.(*common.AppError); ok {
					//appErr.Log = ""
					c.AbortWithStatusJSON(appErr.StatusCode, err)
					panic(err)
					return
				}

				//not app Err
				appErr := common.ErrInternal(err.(error))
				c.AbortWithStatusJSON(appErr.StatusCode, appErr)
				//call panic king again, cause gin have itself recover mechanism
				panic(err)
				return
			}
		}()

		c.Next()
	}
}
