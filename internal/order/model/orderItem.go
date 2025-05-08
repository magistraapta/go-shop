package model

import (
	"gorm.io/gorm"
)

type OrderItem struct {
	gorm.Model
	TransactionID uint    `gorm:"not null"`
	ProductID     uint    `gorm:"not null"`
	Quantity      int     `gorm:"not null"`
	PriceAtTime   float64 `gorm:"not null"`
}
