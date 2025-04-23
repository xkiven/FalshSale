package logic

import (
	kafka2 "FlashSale/dao/kafka"
	"FlashSale/dao/mysql"
	redis2 "FlashSale/dao/redis"
	"FlashSale/kitex_gen/FlashSale/order_service"
	"FlashSale/models"
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/segmentio/kafka-go"
	"gorm.io/gorm"
)

func CreateOrder(ctx context.Context, db *gorm.DB, rdb *redis.Client, producer *kafka.Writer, consumer *kafka.Reader, req *order_service.CreateOrderRequest) (*order_service.CreateOrderResponse, error) {
	order := models.Order{
		UserID:     int(req.UserId),
		ActivityID: int(req.ActivityId),
		ProductID:  int(req.ProductId),
	}

	success, err := redis2.CheckAndDeductStock(rdb)
	{
		if err != nil {
			return &order_service.CreateOrderResponse{
				Code:    1,
				Message: err.Error(),
			}, err
		}
	}
	if !success {
		return &order_service.CreateOrderResponse{
			Code:    4,
			Message: "商品已售空",
		}, err
	}

	topic := "flash_sale"
	err = kafka2.SendOrderMessage(producer, order, topic)
	if err != nil {
		return &order_service.CreateOrderResponse{
			Code:    2,
			Message: err.Error(),
		}, err
	}

	NewOrder, err := kafka2.ConsumeMessage(consumer)
	if err != nil {
		return &order_service.CreateOrderResponse{
			Code:    3,
			Message: err.Error(),
		}, err
	}

	err = mysql.CreateOrder(db, NewOrder)
	if err != nil {
		return &order_service.CreateOrderResponse{
			Code:    1,
			Message: err.Error(),
		}, err
	}

	return &order_service.CreateOrderResponse{
		Code:    0,
		Message: "创建订单成功",
	}, nil
}
