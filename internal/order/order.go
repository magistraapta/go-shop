package order

import (
	cartRepo "golang-shop/internal/cart/repository"
	"golang-shop/internal/order/handler"
	"golang-shop/internal/order/repository"
	"golang-shop/internal/order/services"
	productRepo "golang-shop/internal/product/repository"
	"golang-shop/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupOrder(router *gin.Engine, db *gorm.DB) {
	orderRepo := repository.NewOrderRepository(db)
	productRepo := productRepo.NewProductRepository(db)
	cartRepo := cartRepo.NewCartRepository(db)
	orderService := services.NewOrderServices(orderRepo, productRepo, cartRepo)
	orderHandler := handler.NewOrderHandler(orderService)

	router.POST("api/v1/order/checkout", middleware.RoleMiddleware("user"), orderHandler.Checkout)
}
