package controller

import (
	"net/http"
	"time"
	"orderfood/models"
	"github.com/gin-gonic/gin"
)

var orders = make(map[string]models.Order)

// AddOrder
// @Summary 新增訂單
// @Tags order module
// @Authentication BearerToken
// @Param order body object{restaurant_id=string,order_list=map[string]int} false "訂單資料"
// @Success 200 {string} json{"code", "message"}
// @Router /order/addOrder [post]
// AddOrder
func AddOrder(c *gin.Context) {
    var requestData map[string]interface{}
    if err := c.ShouldBindJSON(&requestData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "Invalid request data", "error": err.Error()})
        return
    }

    order := models.Order{}

    if restaurantId, ok := requestData["restaurant_id"].(string); ok {
        order.RestaurantId = restaurantId
    } else {
        c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "Invalid or missing 'restaurant_id'"})
        return
    }

    if userId, ok := requestData["user_id"].(string); ok {
        order.UserId = userId
    } else {
        c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "Invalid or missing 'user_id'"})
        return
    }

	orderItems := make(map[string]int)
	if orderList, ok := requestData["order_list"].(map[string]interface{}); ok {
		for key, value := range orderList {
			if quantity, ok := value.(float64); ok { 
				orderItems[key] = int(quantity) 
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "Invalid 'order_list' format"})
				return
			}
		}
		order.OrderItems = orderItems
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "Invalid or missing 'order_list'"})
		return
	}
	
    order.ID = time.Now().Format("20060102150405")
    order.OrderTime = time.Now().Format(time.RFC3339)
    order.IsFinished = false

    orders[order.ID] = order

    c.JSON(http.StatusOK, gin.H{"code": 200, "message": "Order added successfully", "order_id": order.ID})
}


// EndOrder
// @Summary 結束訂單
// @Tags order module
// @Authentication BearerToken
// @Param order body object{order_id=string} false "訂單ID"
// @Success 200 {string} json{"code", "message"}
// @Router /order/endOrder [post]
func EndOrder(c *gin.Context) {
	var requestData struct {
		OrderID string `json:"order_id"`
	}
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "Invalid request data"})
		return
	}

	order, exists := orders[requestData.OrderID]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"code": 404, "message": "Order not found"})
		return
	}

	order.IsFinished = true
	orders[requestData.OrderID] = order

	c.JSON(http.StatusOK, gin.H{"code": 200, "message": "Order ended successfully"})
}

// GetOrder
// @Summary 取得訂單
// @Tags order module
// @Authentication BearerToken
// @Success 200 {string} json{"code", "message", "data []Order"}
// @Router /order/getOrder [get]
func GetOrderList(c *gin.Context) {
	var orderList []models.Order
	for _, order := range orders {
		convertedOrder := models.Order{
			ID:           order.ID,
			RestaurantId: order.RestaurantId,
			UserId:       order.UserId,
			OrderItems:   order.OrderItems,
			OrderTime:    order.OrderTime,
			IsFinished:   order.IsFinished,
		}
		orderList = append(orderList, convertedOrder)
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Order list retrieved successfully",
		"data":    orderList,
	})
}

