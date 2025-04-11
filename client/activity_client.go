package client

import (
	"FlashSale/kitex_gen/FlashSale/activity_service"
	"FlashSale/kitex_gen/FlashSale/activity_service/activityservice"
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
)

func GenerateActivityClient(req *activity_service.CreateActivityRequest) (*activity_service.CreateActivityResponse, error) {
	// 创建 Kitex 客户端
	fmt.Println("创建客户端")
	cli, err := activityservice.NewClient(
		"GenerateActivityService",
		client.WithHostPorts("127.0.0.1:8888"),
	)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	resp, err := cli.CreateActivity(context.Background(), req)
	if err != nil {
		panic(err)
	}

	return resp, nil
}
