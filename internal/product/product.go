package product

import (
	"golang-shop/internal/product/handler"
	"golang-shop/internal/product/repository"
	"golang-shop/internal/product/services"
	"golang-shop/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ProductRouter(router *gin.Engine, db *gorm.DB) {
	productRepo := repository.NewProductRepository(db)
	productServices := services.NewProductServices(productRepo)
	productHandler := handler.NewProductHandler(productServices)

	product := router.Group("api/v1/product")
	{
		product.POST("/", middleware.RoleMiddleware("admin"), productHandler.CreateProduct)
		product.GET("/:id", productHandler.GetProductById)
		product.GET("/", productHandler.GetAllProduct)
		product.PUT("/:id", middleware.RoleMiddleware("admin"), productHandler.UpdateProductById)
		product.DELETE("/:id", middleware.RoleMiddleware("admin"), productHandler.DeleteProductById)
	}

}
