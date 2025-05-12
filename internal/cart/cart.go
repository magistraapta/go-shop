package cart

import (
	"golang-shop/internal/cart/handler"
	"golang-shop/internal/cart/repository"
	"golang-shop/internal/cart/services"
	productRepo "golang-shop/internal/product/repository"
	"golang-shop/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupCart(router *gin.Engine, db *gorm.DB) {
	repo := repository.NewCartRepository(db)
	productRepo := productRepo.NewProductRepository(db)
	services := services.NewCartServices(repo, productRepo)
	cartHandler := handler.NewCartHandler(services)

	cart := router.Group("api/v1/cart")
	{
		cart.GET("/", middleware.RoleMiddleware("user"), cartHandler.GetCart)
		cart.POST("/add", middleware.RoleMiddleware("user"), cartHandler.AddToCart)
		cart.DELETE("/item/:id", middleware.RoleMiddleware("user"), cartHandler.RemoveItem)
		cart.PUT("/item/:id", middleware.RoleMiddleware("user"), cartHandler.UpdateQuantity)
	}

}
