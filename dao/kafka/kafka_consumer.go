package kafka

import "github.com/segmentio/kafka-go"

// KafkaConsumer 定义 Kafka 消费者结构体
type KafkaConsumer struct {
	reader *kafka.Reader
}

// NewKafkaConsumer 创建 Kafka 消费者实例
func NewKafkaConsumer(brokers []string, topic string) *kafka.Reader {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		Topic:   topic,
		GroupID: "message-group",
	})
	return reader
}
