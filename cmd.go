package main

import (
	"FlashSale/http_handler"
	"FlashSale/middleware"
	"github.com/gin-gonic/gin"
)

func main() {

	// 创建 Gin 实例
	r := gin.Default()

	r.POST("/activity", http_handler.GenerateActivityHandler)
	// 应用 IP 限流中间件
	r.Use(middleware.IpRateLimit())

	r.POST("/order", http_handler.GenerateOrderHandler)

	// 启动 HTTP 服务

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}

}
