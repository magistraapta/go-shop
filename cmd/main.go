//   Product Api:
//    version: 0.1
//    title: Product Api
//   Schemes: http, https
//   Host:
//   BasePath: /api/v1
//      Consumes:
//      - application/json
//   Produces:
//   - application/json
//   SecurityDefinitions:
//    Bearer:
//     type: apiKey
//     name: Authorization
//     in: header
//   swagger:meta

package main

import (
	"golang-shop/initializers"
	"golang-shop/internal/auth"
	"golang-shop/internal/cart"
	"golang-shop/internal/order"
	"golang-shop/internal/product"
	"log"
	"os"

	_ "golang-shop/docs" // This will be generated

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title E-Commerce API
// @version 1.0
// @description This is a sample E-Commerce API sample.
// @host localhost:8080
// @BasePath /api/v1
// @schemes http https
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
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

	// Swagger documentation route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// module router
	auth.SetupAuth(router, db)
	product.ProductRouter(router, db)
	order.SetupOrder(router, db)
	cart.SetupCart(router, db)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080" // default port
	}
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
}
