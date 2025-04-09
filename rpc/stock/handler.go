package main

import (
	stock_service "FalshSale/kitex_gen/FlashSale/stock_service"
	"context"
)

// StockServiceImpl implements the last service interface defined in the IDL.
type StockServiceImpl struct{}

// DeductStock implements the StockServiceImpl interface.
func (s *StockServiceImpl) DeductStock(ctx context.Context, req *stock_service.DeductStockRequest) (resp *stock_service.DeductStockResponse, err error) {
	// TODO: Your code here...
	return
}
