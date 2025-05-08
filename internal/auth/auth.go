package auth

import (
	"golang-shop/internal/auth/handler"
	"golang-shop/internal/auth/repository"
	"golang-shop/internal/auth/services"
	"golang-shop/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupAuth(router *gin.Engine, db *gorm.DB) {
	authRepo := repository.NewAuthRepository(db)
	authService := services.NewAuthService(authRepo)
	authHandler := handler.NewAuthHandler(authService)

	router.POST("/auth/login", authHandler.Login)

	admin := router.Group("/admin")
	{
		admin.POST("/register", authHandler.RegisterAdmin)
		admin.GET("/dashboard", middleware.RoleMiddleware("admin"), func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "this is admin"})
		})
	}

	user := router.Group("/user")
	{
		user.POST("/register", authHandler.RegisterUser)
	}
}
