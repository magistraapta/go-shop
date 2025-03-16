package router

import (
	"golang-shop/internal/handler"
	"golang-shop/internal/repository"
	"golang-shop/internal/services"

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
		v1.POST("/user", userHandler.CreateUser)
		v1.GET("/user", userHandler.TestApi)
		v1.POST("/login", userHandler.Login)
	}

	return router
}
