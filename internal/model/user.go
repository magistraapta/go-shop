package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username     string        `gorm:"unique;not null"`
	Email        string        `gorm:"unique;not null"`
	Password     string        `gorm:"not null"`
	Cart         Cart          `gorm:"foreignKey:UserID"`
	Transactions []Transaction `gorm:"foreignKey:user_id" json:"transactions"`
}
