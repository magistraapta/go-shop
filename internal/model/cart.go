package model

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	ID     uint       `gorm:"primaryKey" json:"id"`
	UserID uint       `gorm:"uniqueIndex" json:"user_id"`
	Items  []CartItem `json:"items" gorm:"foreignKey:CartID"`
}
