package main

import (
	order_service "FlashSale/kitex_gen/FlashSale/order_service"
	"FlashSale/logic"
	"FlashSale/svc"
	"context"
)

// OrderServiceImpl implements the last service interface defined in the IDL.
type OrderServiceImpl struct {
	sc *svc.ServiceContext
}

// NewOrderServiceImpl 创建 OrderServiceImpl 实例
func NewOrderServiceImpl(sc *svc.ServiceContext) *OrderServiceImpl {
	return &OrderServiceImpl{
		sc: sc,
	}
}

// CreateOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) CreateOrder(ctx context.Context, req *order_service.CreateOrderRequest) (resp *order_service.CreateOrderResponse, err error) {
	// TODO: Your code here...
	// 将单个请求包装成切片
	reqs := []*order_service.CreateOrderRequest{req}
	responses := logic.CreateOrder(ctx, s.sc.MySQLClient, s.sc.RedisClient, s.sc.KafkaProducer, s.sc.KafkaConsumer, reqs)
	if len(responses) > 0 {
		resp = responses[0]
	}
	return
}
