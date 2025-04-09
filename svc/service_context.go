package svc

import (
	"FlashSale/config"
	"FlashSale/dao/kafka"
	"FlashSale/dao/mysql"
	"FlashSale/dao/redis"
)

// ServiceContext 定义服务上下文结构体
type ServiceContext struct {
	Config        config.Config
	KafkaProducer *kafka.KafkaProducer
	MySQLClient   *mysql.MySQLClient
	RedisClient   *redis.RedisClient
	KafkaConsumer *kafka.KafkaConsumer
}

// NewServiceContext 创建服务上下文实例
func NewServiceContext(mysqlClient *mysql.MySQLClient, redisClient *redis.RedisClient, kafkaProducer *kafka.KafkaProducer, kafkaConsumer *kafka.KafkaConsumer) *ServiceContext {
	return &ServiceContext{
		MySQLClient:   mysqlClient,
		RedisClient:   redisClient,
		KafkaProducer: kafkaProducer,
		KafkaConsumer: kafkaConsumer,
	}
}
