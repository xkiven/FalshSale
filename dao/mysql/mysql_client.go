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
func NewMySQLClient(dataSource string) (*MySQLClient, error) {
	db, err := gorm.Open(mysql.Open(dataSource), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &MySQLClient{
		DB: db,
	}, nil
}
func CreateActivity(db *gorm.DB, activity *models.Activity) error {
	err := db.Create(activity).Error
	if err != nil {
		return err
	}
	return nil
}
