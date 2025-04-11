package client

import (
	"FlashSale/kitex_gen/FlashSale/order_service"
	"FlashSale/kitex_gen/FlashSale/order_service/orderservice"
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
)

func GenerateOrderClient(req *order_service.CreateOrderRequest) (*order_service.CreateOrderResponse, error) {
	// 创建 Kitex 客户端
	fmt.Println("创建客户端")
	cli, err := orderservice.NewClient(
		"GenerateOrderService",
		client.WithHostPorts("127.0.0.1:8888"),
	)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	resp, err := cli.CreateOrder(context.Background(), req)
	if err != nil {
		panic(err)
	}

	return resp, nil
}
