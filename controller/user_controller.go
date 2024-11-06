package controller

import (
	"log"
	_ "orderfood/docs"
	"orderfood/models"
	"orderfood/response"
	"orderfood/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// Reigsiter
// @Summary 註冊新用戶
// @Tags user module
// @param user body models.UserBasic false "用戶登入資料"
// @Success 200 {string} json{"code", "message"}
// @Router /user/register [post]

func Register(c *gin.Context) {
	// 使用 c.BindJSON 來從請求中解碼 JSON 到 UserBasic 結構體
	var userInput models.UserBasic
	if err := c.BindJSON(&userInput); err != nil {
		response.Fail(c, nil, "無效的輸入格式")
		return
	}

	// 驗證是否有缺失的字段
	if userInput.Email == "" || userInput.Password == "" || userInput.Name == "" || userInput.Phone == "" || userInput.Roles == "" {
		response.Fail(c, nil, "請輸入完整資料")
		return
	}

	// 加密密碼
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	if err != nil {
		response.Fail(c, nil, "加密密碼失敗")
		return
	}

	// 生成 UUID 並設置到用戶結構體中
	UUID := uuid.New().String()
	userInput.ID = UUID
	userInput.Password = string(hasedPassword)
	userInput.IsDisabled = true // 註冊時設置為默認禁用

	// 將用戶數據存儲到 Firestore
	_, err = utils.FirestoreClient.Collection("users").Doc(UUID).Set(c, userInput)
	if err != nil {
		response.Fail(c, nil, "註冊失敗")
		return
	}

	response.Success(c, nil, "註冊成功")
}

// Login
// @Summary 登入帳號
// @Tags user module
// @param user body models.UserBasic false "用戶登入資料"
// @Success 200 {string} json{"code", "message", "token"}
// @Router /user/login [post]
func Login(c *gin.Context) {
	// 使用 c.BindJSON 來從請求中解析 JSON 到 UserBasic 結構體
	var userInput models.UserBasic
	if err := c.BindJSON(&userInput); err != nil {
		response.Fail(c, nil, "無效的輸入格式")
		return
	}

	// 打印輸入資料
	println("email: ", userInput.Email)
	println("password: ", userInput.Password)

	// 檢查輸入資料
	if userInput.Email == "" || userInput.Password == "" {
		response.Fail(c, nil, "請輸入完整資料")
		return
	}

	// 查詢並獲取文件
	user, err := utils.FindUserByEmail(userInput.Email)

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
		log.Printf("Failed to decode document into User struct: %v", err)
		response.Fail(c, nil, "帳號解析失敗")
		return
	}

	// 驗證密碼
	if err := bcrypt.CompareHashAndPassword([]byte(userBasic.Password), []byte(userInput.Password)); err != nil {
		response.Fail(c, nil, "密碼錯誤")
		return
	}

	// 模擬 token 生成
	token := "11"

	response.Success(c, gin.H{
		"token": token,
	}, "登入成功")
}

// VerifyAccount
// @Summary 驗證帳號
// @Tags user module
// @Param token query string false "token"
// @Success 200 {string} json{"code", "message"}
// @Router /user/verify [post]
func VerifyAccount(c *gin.Context) {
}

// ForgetPassword
// @Summary 忘記密碼
// @Tags user module
// @param email query string false "email信箱"
// @Success 200 {string} json{"code", "message"}
// @Router /user/forget [post]
func ForgetPassword(c *gin.Context) {
}
