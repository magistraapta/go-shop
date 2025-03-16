package initializers

import (
	"golang-shop/internal/model"
	"log"
)

func SyncDatabase() {
	db, err := ConnectDatabase()
	db.AutoMigrate(&model.User{}, &model.Product{}, &model.ShoppingCart{})

	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}
