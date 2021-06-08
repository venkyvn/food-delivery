package main

import (
	"github.com/gin-gonic/gin"
	"go-food-delivery/component"
	"go-food-delivery/component/uploadprovider"
	"go-food-delivery/middleware"
	"go-food-delivery/modules/restaurant/restauranttransport/ginrestaurant"
	"go-food-delivery/modules/restaurantlike/transport/gin"
	"go-food-delivery/modules/upload/uploadtransport/ginupload"
	"go-food-delivery/modules/user/usertransport/ginuser"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

// DBConnectionStr=food_delivery:19e5a718a54a9fe0559dfbce6908@tcp(127.0.0.1:3307)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local
func main() {
	dsn := os.Getenv("DBConnectionStr")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db = db.Debug()

	s3BucketName := os.Getenv("S3BucketName")
	s3Region := os.Getenv("S3Region")
	s3APIKey := os.Getenv("S3APIKey")
	s3SecretKey := os.Getenv("S3SecretKey")
	s3Domain := os.Getenv("S3Domain")
	secretKey := os.Getenv("SecretKey")

	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)

	if err != nil {
		log.Fatalln(err)
	}

	if err := runService(db, s3Provider, secretKey); err != nil {
		log.Fatalln(err)
	}
}

func runService(db *gorm.DB, provider uploadprovider.UploadProvider, secretKey string) error {

	appCtx := component.NewAppContext(db, provider, secretKey)

	r := gin.Default()
	r.Use(middleware.Recover(appCtx))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("/v1")

	v1.POST("/upload", ginupload.Upload(appCtx))
	v1.POST("/register", ginuser.Register(appCtx))
	v1.POST("/login", ginuser.Login(appCtx))
	v1.GET("/profile", middleware.RequireAuth(appCtx), ginuser.GetProfile(appCtx))

	restaurants := v1.Group("/restaurants", middleware.RequireAuth(appCtx))
	{
		restaurants.GET("", ginrestaurant.ListRestaurant(appCtx))
		restaurants.GET("/:id", ginrestaurant.GetRestaurant(appCtx))
		restaurants.POST("", ginrestaurant.CreateRestaurant(appCtx))
		restaurants.PATCH("/:id", ginrestaurant.UpdateRestaurant(appCtx))
		restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appCtx))
		restaurants.GET("/:id/liked-user", ginrestaurantliked.GetUserLikedRestaurant(appCtx))
		restaurants.POST("/:id/like", ginrestaurantliked.UserLikeRestaurant(appCtx))
		restaurants.POST("/:id/unlike", ginrestaurantliked.UserUnlikeRestaurant(appCtx))
	}

	return r.Run()

}
