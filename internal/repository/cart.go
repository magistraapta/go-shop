package repository

import (
	"errors"
	"golang-shop/internal/model"

	"gorm.io/gorm"
)

type CartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{db: db}
}

func (r *CartRepository) FindOrCreateCart(userID uint) (*model.Cart, error) {
	/*
		Input: userID is from user model
		Output: cart Model or error

		args: Find user cart or create new cart if it doesn't exist
	*/
	var cart model.Cart

	result := r.db.Where("user_id = ?", userID).First(&cart)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			newCart := model.Cart{
				UserID: userID,
			}
			if err := r.db.Create(&newCart).Error; err != nil {
				return nil, err
			}

			return &newCart, nil
		}
		return nil, result.Error
	}

	return &cart, nil
}

func (r *CartRepository) AddToCart(cartID uint, item model.CartItem) error {
	item.CartID = cartID
	return r.db.Create(&item).Error
}

func (r *CartRepository) UpdateCartItem(item model.CartItem) error {
	return r.db.Save(&item).Error
}

func (r *CartRepository) GetCartItem(userID uint) (*model.Cart, error) {

	var cart model.Cart

	// find the cart for the user id

	result := r.db.Where("user_id = ?", userID).Preload("Items").First(&cart)

	if result.Error != nil {
		return nil, result.Error
	}

	return &cart, nil
}

func (r *CartRepository) FindCartItem(cartID uint, productID uint) (*model.CartItem, error) {
	var item model.CartItem
	result := r.db.Where("cart_id = ? AND product_id = ?", cartID, productID).First(&item)

	if result.Error != nil {
		return nil, result.Error
	}

	return &item, nil
}

func (r *CartRepository) RemoveCartItem(cartID uint, itemID uint) error {
	return r.db.Where("cart_id = ? AND product_id = ?", cartID, itemID).Delete(&model.CartItem{}).Error
}

func (r *CartRepository) UpdateItemQuantity(productID uint, cartID uint, quantity int) error {
	return r.db.Model(&model.CartItem{}).
		Where("cart_id = ? AND product_id = ?", cartID, productID).
		Update("quantity", quantity).Error
}
