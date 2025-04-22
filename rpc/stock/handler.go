package main

import (
	stock_service "FlashSale/kitex_gen/FlashSale/stock_service"
	"FlashSale/svc"
	"context"
)

// StockServiceImpl implements the last service interface defined in the IDL.
type StockServiceImpl struct {
	sc *svc.ServiceContext
}

// NewStockServiceImpl 创建 StockServiceImpl 实例
func NewStockServiceImpl(sc *svc.ServiceContext) *StockServiceImpl {
	return &StockServiceImpl{
		sc: sc,
	}
}

// DeductStock implements the StockServiceImpl interface.
func (s *StockServiceImpl) DeductStock(ctx context.Context, req *stock_service.DeductStockRequest) (resp *stock_service.DeductStockResponse, err error) {
	// TODO: Your code here...
	return
}
