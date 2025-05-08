package model

import (
	"golang-shop/internal/cart/model"
	transactionModel "golang-shop/internal/order/model"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string                         `gorm:"unique;not null"`
	Email        string                         `gorm:"unique;not null"`
	Password     string                         `gorm:"not null"`
	Role         string                         `gorm:"not null"`
	Cart         model.Cart                     `gorm:"foreignKey:UserID"`
	Transactions []transactionModel.Transaction `gorm:"foreignKey:user_id" json:"transactions"`
}
