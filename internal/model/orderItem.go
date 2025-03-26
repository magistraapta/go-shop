package model

import "time"

type OrderItem struct {
	ID            uint    `gorm:"primaryKey"`
	TransactionID uint    `gorm:"not null"`
	ProductID     uint    `gorm:"not null"`
	Quantity      int     `gorm:"not null"`
	PriceAtTime   float64 `gorm:"not null"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
