package middleware

import (
	"orderfood/models"
	"orderfood/response"
	"orderfood/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenString := c.GetHeader("Authorization")

		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			response.Fail(c, nil, "請登入")
			c.Abort()
			return
		}

		tokenString = tokenString[7:]

		token, claims, err := utils.ParseToken(tokenString)

		if err != nil || !token.Valid {
			response.Fail(c, nil, "請登入")
			c.Abort()
			return
		}

		userID := claims.UserID
		user, err := utils.FindUserByID(userID)

		if err != nil {
			response.Fail(c, nil, "查詢用戶失敗")
			return
		}

		// 檢查用戶是否存在
		if user == nil || !user.Exists() {
			response.Fail(c, nil, "用戶不存在")
			return
		}

		// 使用 User 結構體讀取文件數據
		var userBasic models.UserBasic
		if err := user.DataTo(&userBasic); err != nil {
			response.Fail(c, nil, "帳號解析失敗")
			return
		}

		c.Set("user", userBasic)
	}
}
