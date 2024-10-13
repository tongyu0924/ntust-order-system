package services

/*
導向不同頁面的function的函式都會放在這裡
*/

import (
	"github.com/gin-gonic/gin"
)

// GetIndex
// @Tags 首頁
// @Success 200 {string} Helloworld
// @Router /index [get]
func GetIndex(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}
