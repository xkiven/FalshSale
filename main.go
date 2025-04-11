package main

import (
	"FlashSale/http_handler"
	"github.com/gin-gonic/gin"
)

func main() {

	// 创建 Gin 实例
	r := gin.Default()

	r.POST("/activity", http_handler.GenerateActivityHandler)
	r.POST("/order", http_handler.GenerateOrderHandler)

}
