namespace go
FlashSale.activity_service

//活动创建请求

struct CreateActivityRequest{
    1: i32 activity_id,
    2: string start_time,
    3: string end_time,
    4: i32 product_id,
    5: double discount_price
}

//活动创建响应
struct CreateActivityResponse{
    1: i32 code,
    2: string message
}

//活动服务接口
service ActivityService{
    CreateActivityResponse CreateActivity(
    1: CreateActivityRequest req
    )
}

