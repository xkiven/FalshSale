package http_handler

import (
	"FlashSale/client"
	"FlashSale/kitex_gen/FlashSale/order_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GenerateOrderHandler(c *gin.Context) {
	var req order_service.CreateOrderRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
			"data":  1000,
		})
		return
	}

	// 调用 Kitex 客户端
	resp, err := client.GenerateOrderClient(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
			"data":  1001,
		})
		return
	}
	// 返回 HTTP 响应
	c.JSON(http.StatusOK, gin.H{
		"message": "GenerateOrder successful",
		"data":    resp,
	})
}
