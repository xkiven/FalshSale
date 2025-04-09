package main

import (
	activity_service "FalshSale/kitex_gen/FlashSale/activity_service/activityservice"
	"log"
)

func main() {
	svr := activity_service.NewServer(new(ActivityServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
