package dto

import "golang-shop/internal/product/model"

// Message represents the message response
// @Description Message response
// @Accept json
// @Produce json
// @Param message body string true "Message"
type Message struct {
	Message string `json:"message"`
}

// ProductResponse represents the product response
// @Description Product response
// @Accept json
// @Produce json
// @Param message body string true "Message"
// @Param data body model.Product true "Product"
type ProductResponse struct {
	Message string        `json:"message"`
	Data    model.Product `json:"data"`
}

// ProductRequest represents the product request
// @Description Product request
// @Accept json
// @Produce json
// @Param name body string true "Name"
// @Param stock body int true "Stock"
// @Param price body float64 true "Price"
// @Param description body string true "Description"
type ProductRequest struct {
	Name        string  `json:"name" binding:"required"`
	Stock       int     `json:"stock" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	Description string  `json:"description"`
}

// ProductResponseList represents the product response list
// @Description Product response list
// @Accept json
// @Produce json
// @Param message body string true "Message"
// @Param data body []model.Product true "Product"
type ProductResponseList struct {
	Message string          `json:"message"`
	Data    []model.Product `json:"data"`
}
