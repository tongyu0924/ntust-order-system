package router

import (
	"orderfood/services"

	"github.com/gin-gonic/gin"
)

/*
在此會把所有的路由都設定好，並且將其導向到相對應的function
*/

// SetupRouter sets up the routes for the application
func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.LoadHTMLGlob("views/**/*")
	router.GET("/", services.GetIndex)
	router.GET("/index", services.GetIndex)

	// Example: router.POST("/user/getUserList", services.GetUserList)
	return router
}
