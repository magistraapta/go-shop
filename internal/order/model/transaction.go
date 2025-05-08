package model

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	UserID     uint        `gorm:"not null" json:"user_id"`
	TotalPrice float64     `gorm:"not null" json:"total_price"`
	Payment    Payment     `gorm:"foreignKey:TransactionID" json:"payment"`
	OrderItems []OrderItem `gorm:"foreignKey:TransactionID" json:"order_items"`
}
