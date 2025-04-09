namespace go
FlashSale.stock_service

//扣减库存请求
struct DeductStockRequest{
    1: i32 product_id,
    2: i64 quantity,
}

//扣减库存响应
struct DeductStockResponse{
    1: i32 code,
    2: string message,
    3: bool success,
}

//库存服务
service StockService{
    DeductStockResponse DeductStock(
    1:DeductStockRequest req
    )
}