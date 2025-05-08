package model

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model
	CartID    uint    `json:"cart_id" gorm:"index"`
	ProductID uint    `json:"product_id" gorm:"index"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}
