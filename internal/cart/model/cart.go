package model

// Cart represents the cart model
// @Description Cart model
type Cart struct {
	ID     uint       `gorm:"primarykey" json:"id"`
	UserID uint       `gorm:"uniqueIndex" json:"user_id"`
	Items  []CartItem `json:"items" gorm:"foreignKey:CartID;constraint:OnDelete:CASCADE"`
}
