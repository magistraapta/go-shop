package handler

import (
	"golang-shop/internal/model"
	"golang-shop/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CheckoutHandler struct {
	checkoutService *services.CheckoutService
}

func NewCheckoutHandler(s *services.CheckoutService) *CheckoutHandler {
	return &CheckoutHandler{s}
}

func (h *CheckoutHandler) Checkout(c *gin.Context) {
	// check is user already log in

	user, exist := c.Get("user")

	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "user is unauthorized",
		})
	}

	userID, ok := user.(model.User)

	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Invalid user",
		})
	}

	var request struct {
		PaymentMethod string `json:"payment_method" binding"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid payment method",
		})
	}

	err := h.checkoutService.CheckoutService(userID.ID, request.PaymentMethod)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Checkout successful"})
}
