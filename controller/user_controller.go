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
// @param name body string false "email信箱"
// @param password body string false "密碼"
// @Success 200 {string} json{"code", "message"}
// @Router /user/register [post]
func Register(c *gin.Context) {
	mail := c.PostForm("email")
	password := c.PostForm("password")
	name := c.PostForm("name")
	phone := c.PostForm("phone")
	roles := c.PostForm("roles")
	UUID := uuid.New().String()

	// TODO: some validation
	if mail == "" || password == "" || name == "" || phone == "" || roles == "" {
		response.Fail(c, nil, "請輸入完整資料")
		return
	}

	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Fail(c, nil, "加密密碼失敗")
		return
	}

	_, err = utils.FirestoreClient.Collection("users").Doc(UUID).Set(c, models.UserBasic{
		ID:         UUID,
		Email:      mail,
		Password:   string(hasedPassword),
		Name:       name,
		Phone:      phone,
		IsDisabled: true,
		Roles:      roles,
	})
	if err != nil {
		response.Fail(c, nil, "註冊失敗")
		return
	}
	response.Success(c, nil, "註冊成功")
}

// Login
// @Summary 登入帳號
// @Tags user module
// @param name body string false "email信箱"
// @param password body string false "密碼"
// @Success 200 {string} json{"code", "message", "token"}
// @Router /user/login [post]
func Login(c *gin.Context) {
	// 獲取用戶輸入的資料
	email := c.PostForm("email")
	password := c.PostForm("password")

	// TODO: some validation
	if email == "" || password == "" {
		response.Fail(c, nil, "請輸入完整資料")
		return
	}

	// 查詢並獲取文件
	user, err := utils.FindUserByEmail(email)

	if err != nil {
		response.Fail(c, nil, "查詢用戶失敗")
		return
	}

	// 使用 User 結構體讀取文件數據
	var userBasic models.UserBasic
	if err := user.DataTo(&userBasic); err != nil {
		log.Fatalf("Failed to decode document into User struct: %v", err)
		response.Fail(c, nil, "帳號解析失敗")
	}

	// 驗證密碼
	if err := bcrypt.CompareHashAndPassword([]byte(userBasic.Password), []byte(password)); err != nil {
		response.Fail(c, nil, "密碼錯誤")
		return
	}

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
