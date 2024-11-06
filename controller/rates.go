package controller

import "github.com/gin-gonic/gin"

// AddRate
// @Summary 評價添加/修改
// @Tags rates module
// @Param rate body object{token=string,restaurant_id=string,rate=float32,comment=string} true "評價資料"
// @Success 200 {string} json{"code", "message"}
// @Router /rate/addRate [post]
func AddRate(c *gin.Context) {

}

// DeleteRate
// @Summary 刪除評價
// @Tags rates module
// @Param rate body object{rate_id=string} true "評價ID"
// @Success 200 {string} json{"code", "message"}
// @Router /rate/deleteRate [delete]
func DeleteRate(c *gin.Context) {

}
