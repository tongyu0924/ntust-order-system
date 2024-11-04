package controller

import "github.com/gin-gonic/gin"

// AddRate
// @Summary 評價添加/修改
// @Tags rates module
// @param token body string false "token"
// @param restaurant_id body string false "餐廳ID"
// @param rate body float32 false "評分"
// @param comment body string false "評論"
// @Success 200 {string} json{"code", "message"}
// @Router /rate/addRate [post]
func AddRate(c *gin.Context) {

}

// DeleteRate
// @Summary 刪除評價
// @Tags rates module
// @param token body string false "token"
// @param rate_id body string false "評價ID"
// @Success 200 {string} json{"code", "message"}
// @Router /rate/deleteRate [delete]
func DeleteRate(c *gin.Context) {

}
