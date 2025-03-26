package model

import "time"

type Transaction struct {
	ID         uint    `gorm:"primaryKey"`
	UserID     uint    `gorm:"not null"`
	TotalPrice float64 `gorm:"not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	OrderItems []OrderItem `gorm:"foreignKey:TransactionID"`
}
