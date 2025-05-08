package middleware

import (
	"golang-shop/internal/auth/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RoleMiddleware(role string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get token from cookie
		tokenString, err := ctx.Cookie("Authorization")
		if err != nil || tokenString == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authorization cookie missing"})
			return
		}

		// Validate token
		claims, err := util.ValidateToken(tokenString)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			return
		}

		if role == claims.Role {
			// Set claims in context
			ctx.Set("username", claims.Username)
			ctx.Set("role", claims.Role)

			ctx.Next()
		} else {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "wrong role"})
			return
		}

	}
}
