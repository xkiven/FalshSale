package main

import (
	"FlashSale/http_handler"
	"FlashSale/middleware"
	"github.com/gin-gonic/gin"
)

func main() {

	// 创建 Gin 实例
	r := gin.Default()

	// 应用 IP 限流中间件
	r.Use(middleware.IpRateLimit())

	r.POST("/activity", http_handler.GenerateActivityHandler)
	r.POST("/order", http_handler.GenerateOrderHandler)

}
