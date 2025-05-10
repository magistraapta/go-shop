package model

// CartItem represents the cart item model
// @Description Cart item model
type CartItem struct {
	ID        uint    `gorm:"primarykey" json:"id"`
	CartID    uint    `json:"cart_id" gorm:"index"`
	ProductID uint    `json:"product_id" gorm:"index"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}
