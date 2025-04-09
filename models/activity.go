package models

type Activity struct {
	ID            int     `gorm:"primaryKey;autoIncrement" json:"id"`
	StartTime     string  `gorm:"not null;type:datetime" json:"start_time"`
	EndTime       string  `gorm:"not null;type:datetime" json:"end_time"`
	ProductID     int     `gorm:"not null" json:"product_id"`
	DiscountPrice float64 `gorm:"not null;type:decimal(10, 2)" json:"discount_price"`
	Product       Product `gorm:"foreignKey:ProductID"`
}
