package main

import (
	"FlashSale/config"
	"FlashSale/consul"
	"FlashSale/dao/kafka"
	"FlashSale/dao/mysql"
	"FlashSale/dao/redis"
	activity_service "FlashSale/kitex_gen/FlashSale/activity_service/activityservice"
	"FlashSale/svc"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
)

func main() {
	//加载配置文件
	var cfg config.Config
	err := config.LoadConfig("etc/config.yaml", &cfg)
	if err != nil {
		log.Fatalf("加载配置文件失败: %v", err)
	}
	// 初始化数据库和消息队列客户端
	mysqlClient, err := mysql.NewMySQLClient(cfg.MySQL.DataSource)
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
	// 创建 Consul 注册中心
	consulRegistry, err := consul.NewConsulRegistry("127.0.0.1:8500")
	if err != nil {
		panic(err)
	}

	// 创建服务实例
	svr := activity_service.NewServer(
		activityHandler,
		server.WithRegistry(consulRegistry),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "ActivityServer",
			Tags:        map[string]string{"env": "dev"},
		}),
		server.WithServiceAddr(&net.TCPAddr{ // 设置服务地址和端口
			IP:   net.IPv4(127, 0, 0, 1),
			Port: 8888, // 你的服务端口号
		}),
	)

	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
