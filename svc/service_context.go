package svc

import (
	"FlashSale/config"
	"github.com/go-redis/redis/v8"
	"github.com/segmentio/kafka-go"
	"gorm.io/gorm"
)

// ServiceContext 定义服务上下文结构体
type ServiceContext struct {
	Config        config.Config
	KafkaProducer *kafka.Writer
	MySQLClient   *gorm.DB
	RedisClient   *redis.Client
	KafkaConsumer *kafka.Reader
}

// NewServiceContext 创建服务上下文实例
func NewServiceContext(mysqlClient *gorm.DB, redisClient *redis.Client, kafkaProducer *kafka.Writer, kafkaConsumer *kafka.Reader) *ServiceContext {
	return &ServiceContext{
		MySQLClient:   mysqlClient,
		RedisClient:   redisClient,
		KafkaProducer: kafkaProducer,
		KafkaConsumer: kafkaConsumer,
	}
}
