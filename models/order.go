package models

type Order struct {
	ID          int      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID      int      `gorm:"not null" json:"user_id"`
	ActivityID  int      `gorm:"not null" json:"activity_id"`
	ProductID   int      `gorm:"not null" json:"product_id"`
	OrderStatus string   `gorm:"not null;default:待支付;type:varchar(20)" json:"order_status"`
	Activity    Activity `gorm:"foreignKey:ActivityID"`
	Product     Product  `gorm:"foreignKey:ProductID"`
}
