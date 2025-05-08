package dto

import (
	"golang-shop/internal/cart/model"
	"time"
)

type AddToCartRequest struct {
	ProductID uint `json:"product_id" binding:"required"`
	Quantity  int  `json:"quantity" binding:"required,min=1"`
}

type CartItemsResponse struct {
	CartID    uint             `json:"cart_id"`
	UserID    uint             `json:"user_id"`
	Items     []model.CartItem `json:"items"`
	Total     float64          `json:"total"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
}
