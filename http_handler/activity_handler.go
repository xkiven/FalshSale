package http_handler

import (
	"FlashSale/client"
	"FlashSale/kitex_gen/FlashSale/activity_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GenerateActivityHandler(c *gin.Context) {
	var req activity_service.CreateActivityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
			"data":  1000,
		})
		return
	}
	// 调用 Kitex 客户端
	resp, err := client.GenerateActivityClient(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
			"data":  1001,
		})
		return
	}
	// 返回 HTTP 响应
	c.JSON(http.StatusOK, gin.H{
		"message": "GenerateActivity successful",
		"data":    resp,
	})
}
