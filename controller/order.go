package controller

import "github.com/gin-gonic/gin"

// AddOrder
// @Summary 新增訂單
// @Tags order module
// @Authentication BearerToken
// @param object{restaurant_id=string,order_list=map[string]int} body false "訂單資料"
// @Success 200 {string} json{"code", "message"}
// @Router /order/addOrder [post]
func AddOrder(c *gin.Context) {
}

// EndOrder
// @Summary 結束訂單
// @Tags order module
// @Authentication BearerToken
// @param object{order_id=string} body false "訂單ID"
// @Success 200 {string} json{"code", "message"}
// @Router /order/endOrder [post]
func EndOrder(c *gin.Context) {
}

// GetOrder
// @Summary 取得訂單
// @Tags order module
// @Authentication BearerToken
// @Success 200 {string} json{"code", "message", "data []Order"}
// @Router /order/getOrder [get]
func GetOrderList(c *gin.Context) {
	// 如果角色是商家，則返回該商家的訂單
	// 如果角色是用戶，則返回該用戶的訂單
}
