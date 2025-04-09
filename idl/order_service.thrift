namespace go
FlashSale.order_service

//订单创建请求
struct CreateOrderRequest{
    1: i32 user_id,
    2: i32 activity_id,
    3: i32 product_id,
}

//订单创建响应
struct CreateOrderResponse{
    1: i32 code,
    2: string message,
}

//订单服务接口
service OrderService{
    CreateOrderResponse CreateOrder(
    1: CreateOrderRequest req
    )
}