package model

import "time"

// Product represents the product model
// @Description Product model
type Product struct {
	ID          uint      `gorm:"primarykey" json:"id" example:"1"`
	CreatedAt   time.Time `json:"created_at" example:"2024-01-01T00:00:00Z"`
	UpdatedAt   time.Time `json:"updated_at" example:"2024-01-01T00:00:00Z"`
	DeletedAt   time.Time `gorm:"index" json:"deleted_at,omitempty" example:"2024-01-01T00:00:00Z"`
	Name        string    `gorm:"not null" json:"name" example:"Product 1"`
	Stock       int       `gorm:"not null" json:"stock" example:"10"`
	Price       float64   `gorm:"not null" json:"price" example:"10000"`
	Description string    `json:"description" example:"Description of Product 1"`
}
