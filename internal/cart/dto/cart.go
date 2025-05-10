package dto

import (
	"golang-shop/internal/cart/model"
	"time"
)

// AddToCartRequest represents the add to cart request
// @Description Add to cart request
// @Accept json
// @Produce json
// @Param product_id body uint true "Product ID"
// @Param quantity body int true "Quantity"
type AddToCartRequest struct {
	ProductID uint `json:"product_id" binding:"required"`
	Quantity  int  `json:"quantity" binding:"required,min=1"`
}

// UpdateQuantityRequest represents the update quantity request
// @Description Update quantity request
// @Accept json
// @Produce json
// @Param product_id body uint true "Product ID"
// @Param quantity body int true "Quantity"
type UpdateQuantityRequest struct {
	ProductID uint `json:"product_id" binding:"required"`
	Quantity  int  `json:"quantity" binding:"required,min=1"`
}

// CartItemsResponse represents the cart items response
// @Description Cart items response
// @Accept json
// @Produce json
// @Param cart_id body uint true "Cart ID"
// @Param user_id body uint true "User ID"
type CartItemsResponse struct {
	CartID    uint             `json:"cart_id"`
	UserID    uint             `json:"user_id"`
	Items     []model.CartItem `json:"items"`
	Total     float64          `json:"total"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
}

// CartResponse represents the cart response
// @Description Cart response
// @Accept json
// @Produce json
// @Param message body string true "Message"
type CartResponse struct {
	Message string `json:"message"`
}
