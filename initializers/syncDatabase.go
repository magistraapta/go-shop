package initializers

import (
	"golang-shop/internal/model"
	"log"
)

func SyncDatabase() {
	db, err := ConnectDatabase()
	db.AutoMigrate(
		&model.User{}, &model.Product{},
		&model.Cart{}, &model.CartItem{},
		&model.Transaction{}, &model.OrderItem{},
		&model.Payment{})

	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}
