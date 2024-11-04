package main

import (
	_ "orderfood/docs"
	"orderfood/router"
	"orderfood/utils"
)

func main() {
	r := router.SetupRouter()
	err := utils.InitializeFirebase()

	if err != nil {
		println("Firebase 初始化失敗")
		println(err)
		return
	}

	r.Run(":8080")
}
