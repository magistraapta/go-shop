package initializers

import (
	"golang-shop/internal/model"
	"log"
	"os"

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
			&model.User{}, &model.Product{},
			&model.Cart{}, &model.CartItem{},
			&model.Transaction{}, &model.OrderItem{},
			&model.Payment{},
		)
		if err != nil {
			log.Println("AutoMigrate failed: ", err)
			return nil, err
		}
	}

	log.Println("Connected to database")

	return db, nil
}
