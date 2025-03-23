package router

import (
	"golang-shop/internal/handler"
	"golang-shop/internal/repository"
	"golang-shop/internal/services"
	"golang-shop/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ApiRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	productRepo := repository.NewProductRepo(db)
	productServices := services.NewProductServices(productRepo)
	productHandler := handler.NewProductHandler(productServices)

	cartRepo := repository.NewCartRepository(db)
	cartServices := services.NewCartServices(cartRepo)
	cartHandler := handler.NewCartHandler(cartServices)

	v1 := router.Group("v1")
	{
		v1.GET("/validate", middleware.RequireAuth, middleware.ValidateUser)

		user := v1.Group("/user")
		{
			user.POST("/", userHandler.CreateUser)
			user.GET("/", userHandler.TestApi)
			user.POST("/login", userHandler.Login)
			user.GET("/:id", middleware.RequireAuth, userHandler.GetUserById)
		}

		product := v1.Group("/product")
		{
			product.GET("/:id", productHandler.GetProductById)
			product.POST("/", productHandler.CreateProduct)
			product.DELETE("/:id", productHandler.DeleteProductById)
			product.GET("/", productHandler.GetAllProduct)
			product.PUT("/:id", productHandler.UpdateProductById)
		}

		cart := v1.Group("/cart")
		{
			cart.GET("/", middleware.RequireAuth, cartHandler.GetCart)
			cart.POST("/add", middleware.RequireAuth, cartHandler.AddToCart)
			cart.DELETE("/item/:id", middleware.RequireAuth, cartHandler.RemoveItem)
			cart.PUT("/item/:id", middleware.RequireAuth, cartHandler.UpdateQuantity)
		}

	}

	return router
}
