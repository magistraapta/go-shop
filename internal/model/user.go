package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string       `gorm:"unique;not null"`
	Password string       `gorm:"not null"`
	Email    string       `gorm:"unique;not null"`
	Cart     ShoppingCart `gorm:"foreignKey:UserID"`
}
