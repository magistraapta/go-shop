package initializers

import (
	"log"
	"os"

	userModel "golang-shop/internal/auth/model"
	cartModel "golang-shop/internal/cart/model"
	transactionModel "golang-shop/internal/order/model"
	productModel "golang-shop/internal/product/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() (*gorm.DB, error) {
	var err error

	databaseConfig := os.Getenv("DATABASE_CONFIG")

	if databaseConfig == "" {
		log.Fatal("Database config is not set")
	}

	db, err := gorm.Open(postgres.Open(databaseConfig), &gorm.Config{})

	if err != nil {
		log.Println("Failed Connect to Database ", err)
		return nil, err
	}

	if db != nil {
		err = db.AutoMigrate(
			&userModel.User{}, &productModel.Product{},
			&cartModel.Cart{}, &cartModel.CartItem{},
			&transactionModel.Transaction{}, &transactionModel.OrderItem{},
			&transactionModel.Payment{},
		)
		if err != nil {
			log.Println("AutoMigrate failed: ", err)
			return nil, err
		}
	}

	log.Println("Connected to database")

	return db, nil
}
