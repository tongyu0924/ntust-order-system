package controller

import "github.com/gin-gonic/gin"

// AddRestaurant
// @Summary 新增餐廳
// @Tags restaurant module
// @param token body string false "token"
// @param restaurant_id body string false "餐廳ID"
// @param name body string false "餐廳名稱"
// @param business_hours body string false "營業時間"
// @param address body string false "地址"
// @param phone body string false "電話"
// @Success 200 {string} json{"code", "message"}
// @Router /restaurant/addRestaurant [post]
func AddRestaurant(c *gin.Context) {
}

// AddFavoriteRestaurant
// @Summary 新增最愛餐廳
// @Tags restaurant module
// @param token body string false "token"
// @param restaurant_id body string false "餐廳ID"
// @Success 200 {string} json{"code", "message"}
// @Router /restaurant/addFavoriteRestaurant [post]
func AddFavoriteRestaurant(c *gin.Context) {
}

// FindRestaurantList
// @Summary 查詢餐廳列表
// @Tags restaurant module
// @param name query string false "餐廳名稱"
// @Success 200 {string} json{"code", "message", "data []Restaurant"}
// @Router /restaurant/findRestaurantList [get]
func FindRestaurantList(c *gin.Context) {
}

// FindFavoriteRestaurantList
// @Summary 查詢最愛餐廳列表
// @Tags restaurant module
// @param token body string false "token"
// @Success 200 {string} json{"code", "message", "data []Restaurant"}
// @Router /restaurant/findFavoriteRestaurantList [get]
func FindFavoriteRestaurantList(c *gin.Context) {
}

// FindRestaurantByID
// @Summary 透過ID查詢餐廳
// @Tags restaurant module
// @param id query string false "餐廳ID"
// @Success 200 {string} json{"code", "message", "data Restaurant"}
// @Router /restaurant/findRestaurantByID [get]
func FindRestaurantByID(c *gin.Context) {
}

// FindRestaurantRating
// @Summary 查詢餐廳評分
// @Tags restaurant module
// @param id query string false "餐廳ID"
// @Success 200 {string} json{"code", "message", "data float64", "data rate"}
// @Router /restaurant/findRestaurantRating [get]
func FindRestaurantRating(c *gin.Context) {
}

// UpdateRestaurantInfomation
// @Summary 修改餐廳資訊
// @Tags restaurant module
// @param token body string false "token"
// @param restaurant_id body string false "餐廳ID"
// @param name body string false "餐廳名稱"
// @param business_hours body string false "營業時間"
// @param address body string false "地址"
// @param phone body string false "電話"
// @Success 200 {string} json{"code", "message"}
// @Router /restaurant/updateRestaurantInfomation [post]
func UpdateRestaurantInfomation(c *gin.Context) {
}

// DeleteRestaurant
// @Summary 刪除餐廳
// @Tags restaurant module
// @param token body string false "token"
// @param restaurant_id body string false "餐廳ID"
// @Success 200 {string} json{"code", "message"}
// @Router /restaurant/deleteRestaurant [delete]
func DeleteRestaurant(c *gin.Context) {
}

// RemoveFavoriteRestaurant
// @Summary 刪除最愛餐廳
// @Tags restaurant module
// @param token body string false "token"
// @param restaurant_id body string false "餐廳ID"
// @Success 200 {string} json{"code", "message"}
// @Router /restaurant/removeFavoriteRestaurant [delete]
func RemoveFavoriteRestaurant(c *gin.Context) {
}
