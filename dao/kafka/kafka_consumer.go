package kafka

import (
	"FlashSale/models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
)

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

// ConsumeMessage 消费者读取消息的函数
func ConsumeMessage(consumer *kafka.Reader) (*models.Order, error) {
	// 从Kafka中读取消息
	log.Println("从Kafka中读取消息")
	msg, err := consumer.ReadMessage(context.Background())
	if err != nil {
		return nil, fmt.Errorf("从Kafka读取消息失败: %v", err)
	}

	// 反序列化消息数据为Order结构体
	log.Println("反序列化消息数据为Order结构体")
	var order models.Order
	err = json.Unmarshal(msg.Value, &order)
	if err != nil {
		return nil, fmt.Errorf("反序列化消息失败: %v", err)
	}
	fmt.Println("读取Kafka消息成功")

	return &order, nil
}
