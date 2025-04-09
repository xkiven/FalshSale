package main

import (
	stock_service "FalshSale/kitex_gen/FlashSale/stock_service/stockservice"
	"log"
)

func main() {
	svr := stock_service.NewServer(new(StockServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
