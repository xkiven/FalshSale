package main

func main() {
	////加载配置文件
	//var cfg config.Config
	//err := config.LoadConfig("etc/config.yaml", &cfg)
	//if err != nil {
	//	log.Fatalf("加载配置文件失败: %v", err)
	//}
	//// 初始化数据库和消息队列客户端
	//mysqlClient, err := mysql.NewMySQLClient(cfg.MySQL.DataSource)
	//if err != nil {
	//	log.Fatalf("初始化 MySQL 失败: %v", err)
	//}
	//redisClient := redis.NewRedisClient(cfg.Redis.Host, cfg.Redis.Pass)
	//kafkaProducer := kafka.NewKafkaProducer(cfg.Kafka.Brokers, cfg.Kafka.Topic)
	//kafkaConsumer := kafka.NewKafkaConsumer(cfg.Kafka.Brokers, cfg.Kafka.Topic)
	//
	//// 创建服务上下文
	//sc := svc.NewServiceContext(mysqlClient, redisClient, kafkaProducer, kafkaConsumer)

}
