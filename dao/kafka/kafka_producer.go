package kafka

import (
	"FlashSale/models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"strconv"
)

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

// SendOrderMessage 发送消息到生产者
func SendOrderMessage(producer *kafka.Writer, order models.Order, topic string) error {
	// 将订单结构体序列化为JSON数据
	log.Println("将订单结构体序列化为JSON数据")
	orderData, err := json.Marshal(order)
	if err != nil {
		log.Println(err)
		return err
	}

	keyStr := strconv.Itoa(order.ID)
	key := []byte(keyStr)
	// 创建Kafka消息
	log.Println("创建Kafka消息")
	msg := kafka.Message{
		Key:   key, // 使用订单ID作为消息键
		Value: orderData,
	}

	// 发送消息到Kafka
	log.Println("发送消息到Kafka")
	err = producer.WriteMessages(context.Background(), msg)
	if err != nil {
		log.Println(err)
		return err

	}

	fmt.Printf("订单 %s 已成功发送到 Kafka 主题 %s\n", order.ID, topic)
	return nil
}
