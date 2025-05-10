package handler

import (
	"golang-shop/internal/auth/model"
	"golang-shop/internal/order/dto"
	"golang-shop/internal/order/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderService *services.OrderServices
}

func NewOrderHandler(orderService *services.OrderServices) *OrderHandler {
	return &OrderHandler{orderService: orderService}
}

// @Summary Checkout
// @Description Checkout
// @Accept json
// @Produce json
// @Tags order
// @Param Authorization header string true "Authorization"
// @Param product body dto.CheckoutRequest true "Checkout"
// @Success 200 {object} dto.OrderMessage
// @Failure 401 {object} dto.OrderMessage
// @Failure 500 {object} dto.OrderMessage
// @Router /order/checkout [post]
func (h *OrderHandler) Checkout(c *gin.Context) {

	user, exist := c.Get("user")

	if !exist {
		c.JSON(http.StatusUnauthorized, dto.OrderMessage{
			Message: "user is unauthorized",
		})
	}

	userID, ok := user.(model.User)

	if !ok {
		c.JSON(http.StatusInternalServerError, dto.OrderMessage{
			Message: "Invalid user",
		})
	}

	var request dto.CheckoutRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, dto.OrderMessage{
			Message: "Invalid payment method",
		})
	}

	err := h.orderService.CheckoutService(userID.ID, request.PaymentMethod)

	if err != nil {
		c.JSON(http.StatusBadRequest, dto.OrderMessage{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.OrderMessage{
		Message: "Checkout successful",
	})
}
