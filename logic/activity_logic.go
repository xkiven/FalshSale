package logic

import (
	"FlashSale/dao/mysql"
	redis2 "FlashSale/dao/redis"
	"FlashSale/kitex_gen/FlashSale/activity_service"
	"FlashSale/models"
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"log"
)

func CreateActivity(ctx context.Context, db *gorm.DB, rdb *redis.Client, req *activity_service.CreateActivityRequest) (*activity_service.CreateActivityResponse, error) {
	// 检查产品是否存在
	var product models.Product
	result := db.First(&product, req.ProductId)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return &activity_service.CreateActivityResponse{
				Code:    3,
				Message: "产品不存在",
			}, result.Error
		}
		return &activity_service.CreateActivityResponse{
			Code:    1,
			Message: result.Error.Error(),
		}, result.Error
	}

	activity := models.Activity{
		ID:            int(req.ActivityId),
		StartTime:     req.StartTime,
		EndTime:       req.EndTime,
		ProductID:     int(req.ProductId),
		DiscountPrice: req.DiscountPrice,
		Product:       product,
	}
	err := mysql.CreateActivity(db, &activity)
	if err != nil {
		return &activity_service.CreateActivityResponse{
			Code:    1,
			Message: err.Error(),
		}, err
	}
	//初始化库存
	initialStock := int64(activity.Product.Stock)
	log.Println(initialStock)
	err = redis2.InitializeStock(rdb, activity.ProductID, initialStock)
	if err != nil {
		return &activity_service.CreateActivityResponse{
			Code:    2,
			Message: err.Error(),
		}, err
	}

	return &activity_service.CreateActivityResponse{
		Code:    0,
		Message: "创建活动成功",
	}, nil
}
