package logic

import (
	"FlashSale/dao/mysql"
	"FlashSale/kitex_gen/FlashSale/activity_service"
	"FlashSale/models"
	"context"
	"gorm.io/gorm"
)

func CreateActivity(ctx context.Context, db *gorm.DB, req *activity_service.CreateActivityRequest) (*activity_service.CreateActivityResponse, error) {
	activity := models.Activity{
		ID:            int(req.ActivityId),
		StartTime:     req.StartTime,
		EndTime:       req.EndTime,
		ProductID:     int(req.ProductId),
		DiscountPrice: req.DiscountPrice,
		Product:       models.Product{},
	}
	err := mysql.CreateActivity(db, &activity)
	if err != nil {
		return &activity_service.CreateActivityResponse{
			Code:    1,
			Message: err.Error(),
		}, err
	}
	return &activity_service.CreateActivityResponse{
		Code:    0,
		Message: "创建活动成功",
	}, nil
}
