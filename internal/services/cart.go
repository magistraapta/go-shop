package services

import (
	"errors"
	"golang-shop/internal/dto"
	"golang-shop/internal/model"
	"golang-shop/internal/repository"

	"gorm.io/gorm"
)

type CartServices struct {
	repo *repository.CartRepository
}

func NewCartServices(repo *repository.CartRepository) *CartServices {
	return &CartServices{repo: repo}
}

func (s *CartServices) AddToCart(userID uint, request dto.AddToCartRequest) error {
	// find user cart
	cart, err := s.repo.FindOrCreateCart(userID)

	if err != nil {
		return nil
	}

	// check if item is already on cart
	existingItem, err := s.repo.FindCartItem(cart.ID, request.ProductID)

	// if it's already on cart update the cart and add the quantity
	if existingItem != nil && err == nil {
		existingItem.Quantity += request.Quantity
		return s.repo.UpdateCartItem(*existingItem)
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if cart.ID == 0 {
		return errors.New("invalid cart ID: cart was not properly created")
	}

	// if not already in user's cart add item to cart
	newItem := model.CartItem{
		CartID:    cart.ID,
		ProductID: request.ProductID,
		Quantity:  request.Quantity,
		Price:     request.Price,
	}

	return s.repo.AddToCart(cart.ID, newItem)

}

func (s *CartServices) GetCartItems(userID uint) (*dto.CartItemsResponse, error) {
	// get cart with items

	cart, err := s.repo.GetCartItem(userID)

	if err != nil {
		return nil, err
	}

	var total float64
	for _, item := range cart.Items {
		total += float64(item.Quantity) * item.Price
	}

	response := dto.CartItemsResponse{
		CartID: cart.ID,
		UserID: userID,
		Items:  cart.Items,
		Total:  total,
	}

	return &response, nil

}

func (r *CartServices) RemoveItem(userID uint, itemID uint) error {

	cart, err := r.repo.GetCartItem(userID)

	if err != nil {
		return err
	}

	return r.repo.RemoveCartItem(cart.ID, itemID)
}

func (s *CartServices) UpdateQuantity(userID, productID uint, quantity int) error {
	cart, err := s.repo.GetCartItem(userID)

	if err != nil {
		return err
	}

	if quantity <= 0 {
		existingItem, err := s.repo.FindCartItem(cart.ID, productID)

		if err != nil {
			return err
		}

		return s.repo.RemoveCartItem(cart.ID, existingItem.CartID)
	}

	return s.repo.UpdateItemQuantity(productID, cart.ID, quantity)
}

func (s *CartServices) GetCartByUserId(userID uint) (model.Cart, error) {
	return s.repo.GetCartByUserID(userID)
}
