package main

import (
	"orderfood/router"
)

func main() {
	r := router.SetupRouter()
	r.Run(":8080")
}
