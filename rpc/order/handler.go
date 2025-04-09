package main

import (
	order_service "FalshSale/kitex_gen/FlashSale/order_service"
	"context"
)

// OrderServiceImpl implements the last service interface defined in the IDL.
type OrderServiceImpl struct{}

// CreateOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) CreateOrder(ctx context.Context, req *order_service.CreateOrderRequest) (resp *order_service.CreateOrderResponse, err error) {
	// TODO: Your code here...
	return
}
