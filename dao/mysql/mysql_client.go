package mysql

import (
	"FlashSale/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// MySQLClient 定义 MySQL 客户端结构体
type MySQLClient struct {
	DB *gorm.DB
}

// NewMySQLClient 创建 MySQL 实例
func NewMySQLClient(dataSource string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dataSource), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// 自动迁移模型
	err = db.AutoMigrate(&models.Activity{}, &models.Order{}, &models.Product{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

// CreateActivity 创建活动
func CreateActivity(db *gorm.DB, activity *models.Activity) error {
	err := db.Create(activity).Error
	if err != nil {
		return err
	}
	return nil
}

// CreateOrder 创建订单
func CreateOrder(db *gorm.DB, order *models.Order) error {
	err := db.Create(order).Error
	if err != nil {
		return err
	}
	return nil
}
