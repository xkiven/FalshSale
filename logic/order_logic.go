package logic

import (
	"FlashSale/dao/mysql"
	redis2 "FlashSale/dao/redis"
	"FlashSale/kitex_gen/FlashSale/order_service"
	"FlashSale/models"
	"context"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

func CreateOrder(ctx context.Context, db *gorm.DB, rdb *redis.Client, req *order_service.CreateOrderRequest) (*order_service.CreateOrderResponse, error) {
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
			Code:    1,
			Message: "商品已售空",
		}, err
	}

	err = mysql.CreateOrder(db, &order)
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
