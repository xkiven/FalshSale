package main

import (
	order_service "FlashSale/kitex_gen/FlashSale/order_service/orderservice"
	"log"
)

func main() {
	svr := order_service.NewServer(new(OrderServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
