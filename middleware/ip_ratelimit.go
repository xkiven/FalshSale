package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"net/http"
	"time"
)

// 初始化缓存
var ipCache = cache.New(5*time.Minute, 10*time.Minute)

// IpRateLimit IP 限流中间件
func IpRateLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		key := "ip:" + ip

		// 获取 IP 请求计数
		count, found := ipCache.Get(key)
		if found {
			// 增加计数
			newCount := count.(int) + 1
			if newCount > 1 { // 每秒最多 1 个请求
				c.JSON(http.StatusTooManyRequests, gin.H{
					"error": "请求过于频繁，请稍后再试",
				})
				c.Abort()
				return
			}
			ipCache.Set(key, newCount, cache.DefaultExpiration)
		} else {
			// 首次请求，设置计数为 1
			ipCache.Set(key, 1, cache.DefaultExpiration)
		}

		c.Next()
	}
}
