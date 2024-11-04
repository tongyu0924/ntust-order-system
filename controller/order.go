package controller

import "github.com/gin-gonic/gin"

// AddOrder
// @Summary 新增訂單
// @Tags order module
// @param token body string false "token"
// @param restaurant_id body string false "餐廳ID"
// @param order_list body map[string]int false "訂單列表"
// @Success 200 {string} json{"code", "message"}
// @Router /order/addOrder [post]
func AddOrder(c *gin.Context) {
}

// EndOrder
// @Summary 結束訂單
// @Tags order module
// @param token body string false "token"
// @param order_id body string false "訂單ID"
// @Success 200 {string} json{"code", "message"}
// @Router /order/endOrder [post]
func EndOrder(c *gin.Context) {
}

// GetOrder
// @Summary 取得訂單
// @Tags order module
// @param token query string false "token"
// @Success 200 {string} json{"code", "message", "data []Order"}
// @Router /order/getOrder [get]
func GetOrderList(c *gin.Context) {

}
