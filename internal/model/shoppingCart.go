package model

import "gorm.io/gorm"

type ShoppingCart struct {
	gorm.Model
	UserID   uint      `gorm:"uniqueIndex"`
	Products []Product `gorm:"many2many:cart_products;"`
}
