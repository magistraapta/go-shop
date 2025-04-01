package model

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserID uint       `gorm:"uniqueIndex" json:"user_id"`
	Items  []CartItem `json:"items" gorm:"foreignKey:CartID;constraint:OnDelete:CASCADE"`
}
