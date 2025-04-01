package model

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	PaymentMethod string `json:"payment_type"`
	Amount        int
	TransactionID uint `gorm:"uniqueIndex"`
}
