package models

type Product struct {
	ID           int     `gorm:"primaryKey;autoIncrement" json:"id"`
	Name         string  `gorm:"not null;type:varchar(255)" json:"name"`
	Price        float64 `gorm:"not null;type:decimal(10, 2)" json:"price"`
	MerchantInfo string  `gorm:"type:text" json:"merchant_info"`
	Stock        int     `gorm:"not null" json:"stock"`
}
