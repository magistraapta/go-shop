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

	}

	return router
}
