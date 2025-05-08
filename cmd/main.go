package main

import (
	"golang-shop/initializers"
	"golang-shop/internal/auth"
	"golang-shop/internal/cart"
	"golang-shop/internal/order"
	"golang-shop/internal/product"
	"log"

	"github.com/gin-gonic/gin"
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

	router := gin.Default()
	auth.SetupAuth(router, db)
	product.ProductRouter(router, db)
	order.SetupOrder(router, db)
	cart.SetupCart(router, db)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
}
