package router

import (
	"orderfood/controller"
	_ "orderfood/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

/*
在此會把所有的路由都設定好，並且將其導向到相對應的function
*/

// SetupRouter sets up the routes for the application
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// 設定 cors
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "DELETE", "OPTIONS", "PUT"}
	config.AllowHeaders = []string{"Authorization", "Content-Type", "Upgrade", "Origin",
		"Connection", "Accept-Encoding", "Accept-Language", "Host", "Access-Control-Request-Method", "Access-Control-Request-Headers"}

	router.Use(cors.New(config))

	router.LoadHTMLGlob("views/**/*")
	router.GET("/", controller.GetIndex)
	router.GET("/index", controller.GetIndex)

	// User
	router.POST("/user/register", controller.Register)
	router.POST("/user/login", controller.Login)
	router.POST("/user/verify", controller.VerifyAccount)
	router.POST("/user/forget", controller.ForgetPassword)

	// Menu
	router.DELETE("/menu/deleteFood", controller.DeleteFood)
	router.POST("/menu/addMenu", controller.AddMenu)

	// Order
	router.POST("/order/addOrder", controller.AddOrder)
	router.POST("/order/endOrder", controller.EndOrder)
	router.GET("/order/getOrder", controller.GetOrderList)

	// Rates
	router.POST("/rate/addRate", controller.AddRate)
	router.DELETE("/rate/deleteRate", controller.DeleteRate)

	// Restaurant
	router.POST("/restaurant/addRestaurant", controller.AddRestaurant)
	router.POST("/restaurant/addFavoriteRestaurant", controller.AddFavoriteRestaurant)
	router.GET("/restaurant/findRestaurantList", controller.FindRestaurantList)
	router.GET("/restaurant/findFavoriteRestaurantList", controller.FindFavoriteRestaurantList)
	router.GET("/restaurant/findRestaurantByID", controller.FindRestaurantByID)
	router.GET("/restaurant/findRestaurantRating", controller.FindRestaurantRating)
	router.POST("/restaurant/updateRestaurantInfomation", controller.UpdateRestaurantInfomation)
	router.DELETE("/restaurant/deleteRestaurant", controller.DeleteRestaurant)
	router.DELETE("/restaurant/removeFavoriteRestaurant", controller.RemoveFavoriteRestaurant)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Example: router.POST("/user/getUserList", services.GetUserList)
	return router
}
