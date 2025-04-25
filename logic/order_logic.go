package logic

import (
	kafka2 "FlashSale/dao/kafka"
	"FlashSale/dao/mysql"
	redis2 "FlashSale/dao/redis"
	"FlashSale/kitex_gen/FlashSale/order_service"
	"FlashSale/models"
	"context"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/go-redis/redis/v8"
	"github.com/segmentio/kafka-go"
	"gorm.io/gorm"
	"log"
)

// OrderTask 定义一个订单处理任务的结构体
type OrderTask struct {
	Ctx      context.Context
	DB       *gorm.DB
	RDB      *redis.Client
	Producer *kafka.Writer
	Consumer *kafka.Reader
	Req      *order_service.CreateOrderRequest
	Result   chan *order_service.CreateOrderResponse
}

// 工作协程函数，处理单个订单任务
func worker(tasks <-chan OrderTask) {
	for task := range tasks {
		resp, err := createSingleOrder(task.Ctx, task.DB, task.RDB, task.Producer, task.Consumer, task.Req)
		if err != nil {
			log.Printf("创建订单失败:%v\n", err)
		}
		task.Result <- resp
	}
}

// 封装原有的 CreateOrder 逻辑到一个新的函数中
func createSingleOrder(ctx context.Context, db *gorm.DB, rdb *redis.Client, producer *kafka.Writer, consumer *kafka.Reader, req *order_service.CreateOrderRequest) (*order_service.CreateOrderResponse, error) {
	var resp *order_service.CreateOrderResponse
	err := hystrix.Do("create_order", func() error {
		order := models.Order{
			UserID:     int(req.UserId),
			ActivityID: int(req.ActivityId),
			ProductID:  int(req.ProductId),
		}

		success, err := redis2.CheckAndDeductStock(rdb, order.ProductID)
		if err != nil {
			resp = &order_service.CreateOrderResponse{
				Code:    1,
				Message: err.Error(),
			}
			return err
		}
		if !success {
			resp = &order_service.CreateOrderResponse{
				Code:    4,
				Message: "商品已售空",
			}
			return nil
		}

		topic := "flash_sale"
		err = kafka2.SendOrderMessage(producer, order, topic)
		if err != nil {
			resp = &order_service.CreateOrderResponse{
				Code:    2,
				Message: err.Error(),
			}
			return err
		}

		NewOrder, err := kafka2.ConsumeMessage(consumer)
		if err != nil {
			resp = &order_service.CreateOrderResponse{
				Code:    3,
				Message: err.Error(),
			}
			return err
		}

		err = mysql.CreateOrder(db, NewOrder)
		if err != nil {
			resp = &order_service.CreateOrderResponse{
				Code:    1,
				Message: err.Error(),
			}
			return err
		}

		resp = &order_service.CreateOrderResponse{
			Code:    0,
			Message: "创建订单成功",
		}
		return nil
	}, func(err error) error {
		// 降级逻辑
		resp = &order_service.CreateOrderResponse{
			Code:    5,
			Message: "系统繁忙，请稍后重试",
		}
		return nil
	})

	return resp, err
}

// CreateOrder 并发处理多个订单的函数
func CreateOrder(ctx context.Context, db *gorm.DB, rdb *redis.Client, producer *kafka.Writer, consumer *kafka.Reader, reqs []*order_service.CreateOrderRequest) []*order_service.CreateOrderResponse {
	// 定义工作协程的数量
	numWorkers := 5
	// 创建任务通道
	tasks := make(chan OrderTask, len(reqs))
	// 创建结果通道
	results := make(chan *order_service.CreateOrderResponse, len(reqs))

	// 启动工作协程
	for i := 0; i < numWorkers; i++ {
		go worker(tasks)
	}

	// 分发任务到任务通道
	for _, req := range reqs {
		tasks <- OrderTask{
			Ctx:      ctx,
			DB:       db,
			RDB:      rdb,
			Producer: producer,
			Consumer: consumer,
			Req:      req,
			Result:   results,
		}
	}
	close(tasks)

	// 收集结果
	var responses []*order_service.CreateOrderResponse
	for i := 0; i < len(reqs); i++ {
		responses = append(responses, <-results)
	}
	close(results)

	return responses
}
