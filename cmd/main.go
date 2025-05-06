package main

import (
	"golang-shop/initializers"
	"golang-shop/internal/router"
	"log"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDatabase()
}

func main() {
	db, err := initializers.ConnectDatabase()

	if err != nil {
		log.Println("Failed to connect database")
	}

	router := router.ApiRouter(db)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
}
