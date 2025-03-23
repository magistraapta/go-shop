package dto

import (
	"golang-shop/internal/model"
	"time"
)

type AddToCartRequest struct {
	ProductID uint    `json:"product_id" binding:"required"`
	Quantity  int     `json:"quantity" binding:"required,min=1"`
	Price     float64 `json:"price" binding:"required,min=0"`
}

type CartItemsResponse struct {
	CartID    uint             `json:"cart_id"`
	UserID    uint             `json:"user_id"`
	Items     []model.CartItem `json:"items"`
	Total     float64          `json:"total"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
}
