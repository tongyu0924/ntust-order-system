package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Response(c *gin.Context, httpCode, errCode int, data interface{}, msg string) {
	c.JSON(httpCode, gin.H{
		"code": errCode,
		"msg":  msg,
		"data": data,
	})
}

func Success(c *gin.Context, data interface{}, msg string) {
	Response(c, http.StatusOK, 200, data, msg)
}

func Fail(c *gin.Context, data interface{}, msg string) {
	Response(c, http.StatusOK, 400, data, msg)
}
