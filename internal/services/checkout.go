package services

import (
	"errors"
	"golang-shop/internal/model"
	"golang-shop/internal/repository"
)

type CheckoutService struct {
	cartRepo        *repository.CartRepository
	productRepo     *repository.ProductRepository
	transactionRepo *repository.TransactionRepository
}

func NewCheckoutServices(cartRepo *repository.CartRepository, productRepo *repository.ProductRepository, transactionRepo *repository.TransactionRepository) *CheckoutService {
	return &CheckoutService{cartRepo: cartRepo, productRepo: productRepo, transactionRepo: transactionRepo}
}

func (s *CheckoutService) CheckoutService(userID uint) error {
	cart, err := s.cartRepo.GetCartByUserID(userID)

	if err != nil {
		return errors.New("cart not found")
	}

	if len(cart.Items) == 0 {
		return errors.New("Cart is empty")
	}

	var totalPrice float64
	var orderItems []model.OrderItem

	for _, item := range orderItems {
		err := s.productRepo.UpdateProductStock(item.ProductID, item.Quantity)
		if err != nil {
			return errors.New("Failed to update product stock")
		}

		totalPrice += float64(item.Quantity) * item.PriceAtTime

		orderItems = append(orderItems, model.OrderItem{
			ProductID:   item.ProductID,
			Quantity:    item.Quantity,
			PriceAtTime: item.PriceAtTime,
		})
	}

	transaction := model.Transaction{
		UserID:     cart.UserID,
		TotalPrice: totalPrice,
	}

	err = s.transactionRepo.CreateTransaction(transaction, orderItems)

	if err != nil {
		return errors.New("failed to add items to cart")
	}

	err = s.cartRepo.ClearCart(cart.ID)
	if err != nil {
		return errors.New("failed to empty cart")
	}

	return nil
}
