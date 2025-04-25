package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"sync"
)

// RedisClient 定义 Redis 客户端结构体
type RedisClient struct {
	Client *redis.Client
}

// NewRedisClient 创建 Redis 客户端实例
func NewRedisClient(host, pass string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: pass,
		DB:       0,
	})
	return client
}

var ctx = context.Background()
var mu sync.Mutex

// 假设商品在 Redis 中的 key 为 "product:stock:{productID}"，这里以 productID 为 1 举例
const productStockKey = "product:stock:1"

// InitializeStock 函数用于初始化商品库存
func InitializeStock(client *redis.Client, initialStock int64) error {
	mu.Lock()         //加锁
	defer mu.Unlock() //函数结束时解锁
	setResult := client.Set(ctx, productStockKey, initialStock, 0)
	return setResult.Err()
}

// CheckAndDeductStock 函数用于检查和扣减库存
func CheckAndDeductStock(client *redis.Client, productID int) (bool, error) {
	// 动态生成库存键
	productStockKey := fmt.Sprintf("product:stock:%d", productID)

	// 从 Redis 中获取当前库存数量并原子性扣减
	decrementResult := client.Decr(ctx, productStockKey)
	if decrementResult.Err() != nil {
		log.Printf("Redis 扣减库存失败，商品 ID: %d，错误信息: %v", productID, decrementResult.Err())
		return false, decrementResult.Err()
	}

	// 获取扣减后的库存数量
	stock := decrementResult.Val()
	if stock >= 0 {
		log.Printf("秒杀成功，商品 ID: %d，继续后续订单创建流程", productID)
		return true, nil
	}

	log.Printf("秒杀失败，商品 ID: %d，商品已售罄", productID)
	// 库存不足，将库存加回去
	_, err := client.Incr(ctx, productStockKey).Result()
	if err != nil {
		log.Printf("Redis 恢复库存失败，商品 ID: %d，错误信息: %v", productID, err)
	}
	return false, nil
}
