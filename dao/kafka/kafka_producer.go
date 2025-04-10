package kafka

import "github.com/segmentio/kafka-go"

// KafkaProducer 定义 Kafka 生产者结构体
type KafkaProducer struct {
	writer *kafka.Writer
}

// NewKafkaProducer 创建 Kafka 生产者实例
func NewKafkaProducer(brokers []string, topic string) *kafka.Writer {
	writer := &kafka.Writer{
		Addr:     kafka.TCP(brokers...),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
	return writer
}
