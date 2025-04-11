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
	resp, err = logic.CreateOrder(ctx, s.sc.MySQLClient, s.sc.RedisClient, req)
	return
}
