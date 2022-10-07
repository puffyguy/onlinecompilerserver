package main

import (
	"onlinecompiler/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/getCode", controller.GetCode)
	router.Run(":8000")
}
