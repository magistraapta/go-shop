package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidateUser(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"message": "Im logged in!",
		"user":    user,
	})
}
