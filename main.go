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
	"go-food-delivery/pubsub/pubsublocal"
	"go-food-delivery/skio"
	"go-food-delivery/subscriber"
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

	pubSub := pubsublocal.NewLocalPubSub()
	appCtx := component.NewAppContext(db, provider, secretKey, pubSub)
	r := gin.Default()

	rtEngine := skio.NewRtEngine()

	if err := rtEngine.Run(appCtx, r); err != nil {
		log.Fatal(err)
	}

	//subscriber.Setup(appCtx)
	if err := subscriber.NewEngine(appCtx, rtEngine).Start(); err != nil {
		log.Fatal(err)
	}

	r.Use(middleware.Recover(appCtx))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.StaticFile("/demo/", "./demo.html")
	r.StaticFile("/demo2/", "./demo2.html")

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

	//startSocketIOServer(r, appCtx)

	return r.Run()
}

//func startSocketIOServer(engine *gin.Engine, appCtx component.AppContext) {
//	server, _ := socketio.NewServer(&engineio.Options{
//		Transports: []transport.Transport{websocket.Default},
//	})
//
//	server.OnConnect("/", func(s socketio.Conn) error {
//		//s.SetContext("")
//		fmt.Println("connected:", s.ID(), " IP:", s.RemoteAddr())
//
//		go func() {
//			var i = 0
//			for {
//				i++
//				s.Emit("test",i)
//				time.Sleep(time.Second)
//			}
//		}()
//
//		//s.Join("Shipper")
//		//server.BroadcastToRoom("/", "Shipper", "test", "Hello 200lab")
//
//		return nil
//	})
//
//	server.OnError("/", func(s socketio.Conn, e error) {
//		fmt.Println("meet error:", e)
//	})
//
//	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
//		fmt.Println("closed", reason)
//		// Remove socket from socket engine (from app context)
//	})
//
//	server.OnEvent("/", "authenticate", func(s socketio.Conn, token string) {
//
//		// Validate token
//		// If false: s.Close(), and return
//
//		// If true
//		// => UserId
//		// Fetch db find user by Id
//		// Here: s belongs to who? (user_id)
//		// We need a map[user_id][]socketio.Conn
//
//		db := appCtx.GetMainDBConnection()
//		store := userstorage.NewSqlStore(db)
//		//
//		tokenProvider := jwt.NewJwtProvider(appCtx.GetSecretKey())
//		//
//		payload, err := tokenProvider.Validate(token)
//
//		if err != nil {
//			s.Emit("authentication_failed", err.Error())
//			s.Close()
//			return
//		}
//		//
//		user, err := store.FindUser(context.Background(), map[string]interface{}{"id": payload.UserId})
//		//
//		if err != nil {
//			s.Emit("authentication_failed", err.Error())
//			s.Close()
//			return
//		}
//
//		if user.Status == 0 {
//			s.Emit("authentication_failed", errors.New("you has been banned/deleted"))
//			s.Close()
//			return
//		}
//
//		user.Mask(false)
//
//		s.Emit("your_profile", user)
//		s.Emit("test","data test")
//
//
//
//	})
//
//	server.OnEvent("/", "test", func(s socketio.Conn, msg string) {
//		log.Println(msg)
//	})
//
//	type Person struct {
//		Name string `json:"name"`
//		Age  int    `json:"age"`
//	}
//
//	server.OnEvent("/", "notice", func(s socketio.Conn, p Person) {
//		fmt.Println("server receive notice:", p.Name, p.Age)
//
//		p.Age = 33
//		s.Emit("notice", p)
//	})
//
//	server.OnEvent("/", "test", func(s socketio.Conn, msg string) {
//		fmt.Println("server receive test:", msg)
//	})
//	//
//	//server.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) string {
//	//	s.SetContext(msg)
//	//	return "recv " + msg
//	//})
//	//
//	//server.OnEvent("/", "bye", func(s socketio.Conn) string {
//	//	last := s.Context().(string)
//	//	s.Emit("bye", last)
//	//	s.Close()
//	//	return last
//	//})
//	//
//	//server.OnEvent("/", "noteSumit", func(s socketio.Conn) string {
//	//	last := s.Context().(string)
//	//	s.Emit("bye", last)
//	//	s.Close()
//	//	return last
//	//})
//
//	go server.Serve()
//
//	engine.GET("/socket.io/*any", gin.WrapH(server))
//	engine.POST("/socket.io/*any", gin.WrapH(server))
//}
