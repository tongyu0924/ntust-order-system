package controller

import (
	"net/http"
	"orderfood/models"
	"github.com/gin-gonic/gin"
)

var rates = make(map[string]models.Rate)

// AddRate
// @Summary 評價添加/修改
// @Tags rates module
// @Param rate body object{restaurant_id=string, user_id=string, rate=int, comment=string} true "評價資料"
// @Success 200 {string} json{"code", "message"}
// @Router /rate/addRate [post]
func AddRate(c *gin.Context) {
	var requestData map[string]interface{}
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "Invalid request data", "error": err.Error()})
		return
	}

	rate := models.Rate{}

	if restaurantId, ok := requestData["restaurant_id"].(string); ok {
		rate.RestaurantId = restaurantId
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "Invalid or missing 'restaurant_id'"})
		return
	}

	if userId, ok := requestData["user_id"].(string); ok {
		rate.UserId = userId
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "Invalid or missing 'user_id'"})
		return
	}

	if rateValue, ok := requestData["rate"].(float64); ok { 
		rate.Rate = int(rateValue)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "Invalid or missing 'rate'"})
		return
	}

	if comment, ok := requestData["comment"].(string); ok {
		rate.Comment = comment
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "Invalid or missing 'comment'"})
		return
	}

	key := rate.RestaurantId + "_" + rate.UserId
	rates[key] = rate

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Rate added/updated successfully",
		"key":     key,
		"data":    rate,
	})
}


// DeleteRate
// @Summary 刪除評價
// @Tags rates module
// @Param rate body object{restaurant_id=string, user_id=string} true "評價資料"
// @Success 200 {string} json{"code", "message"}
// @Router /rate/deleteRate [delete]
func DeleteRate(c *gin.Context) {
	var requestData struct {
		RestaurantId string `json:"restaurant_id"`
		UserId       string `json:"user_id"`
	}
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "Invalid request data", "error": err.Error()})
		return
	}

	key := requestData.RestaurantId + "_" + requestData.UserId
	if _, exists := rates[key]; !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "Rate not found",
			"key":     key,
		})
		return
	}

	delete(rates, key)
	c.JSON(http.StatusOK, gin.H{"code": 200, "message": "Rate deleted successfully"})
}
