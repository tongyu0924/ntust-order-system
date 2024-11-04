package controller

import (
	_ "orderfood/models"

	"github.com/gin-gonic/gin"
)

// AddMenu
// @Summary 新增/修改菜單
// @Tags menu module
// @Param token body string false "token"
// @Param restaurant_id body string false "餐廳ID"
// @Param id body string true "Food ID"
// @Success 200 {string} json{"code", "message"}
// @Router /food/addMenu [post]
func AddMenu(c *gin.Context) {

}

// DeleteFood
// @Summary 刪除食物
// @Tags menu module
// @param token body string false "token"
// @param food_name body string false "食物名稱"
// @Success 200 {string} json{"code", "message"}
// @Router /menu/deleteFood [delete]
func DeleteFood(c *gin.Context) {

}
