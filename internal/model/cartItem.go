package model

type CartItem struct {
	ID        uint    `json:"id" gorm:"foreignKey"`
	CartID    uint    `json:"cart_id"`
	ProductID uint    `json:"product_id" gorm:"index"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}
