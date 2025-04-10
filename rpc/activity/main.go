package main

import (
	"FlashSale/config"
	"FlashSale/dao/kafka"
	"FlashSale/dao/redis"
	activity_service "FlashSale/kitex_gen/FlashSale/activity_service/activityservice"
	"FlashSale/svc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	//加载配置文件
	var cfg config.Config
	err := config.LoadConfig("etc/config.yaml", &cfg)
	if err != nil {
		log.Fatalf("加载配置文件失败: %v", err)
	}
	log.Println(cfg)
	// 初始化数据库和消息队列客户端
	url := cfg.MySQL.DataSource
	log.Println(url)
	mysqlClient, err := gorm.Open(mysql.Open(cfg.MySQL.DataSource), &gorm.Config{})
	if err != nil {
		log.Fatalf("初始化 MySQL 失败: %v", err)
	}
	redisClient := redis.NewRedisClient(cfg.Redis.Host, cfg.Redis.Pass)
	kafkaProducer := kafka.NewKafkaProducer(cfg.Kafka.Brokers, cfg.Kafka.Topic)
	kafkaConsumer := kafka.NewKafkaConsumer(cfg.Kafka.Brokers, cfg.Kafka.Topic)
	// 创建服务上下文
	sc := svc.NewServiceContext(mysqlClient, redisClient, kafkaProducer, kafkaConsumer)

	// 初始化 ActivityService
	activityHandler := NewActivityServiceImpl(sc)
	svr := activity_service.NewServer(activityHandler)

	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
