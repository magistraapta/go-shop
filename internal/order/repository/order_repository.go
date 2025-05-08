package repository

import (
	"errors"
	cartModel "golang-shop/internal/cart/model"
	"golang-shop/internal/order/model"

	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) CreateOrder(transaction model.Transaction, items []model.OrderItem, paymentMethod string) error {
	if err := r.db.Create(&transaction).Error; err != nil {
		return err

	}

	for i := range items {
		items[i].TransactionID = transaction.ID
		if err := r.db.Create(&items[i]).Error; err != nil {
			return err
		}
	}

	payment := model.Payment{
		TransactionID: transaction.ID,
		PaymentMethod: paymentMethod,
		Amount:        int(transaction.TotalPrice),
	}

	if err := r.db.Create(&payment).Error; err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) FindOrCreateCart(userID uint) (*cartModel.Cart, error) {
	/*
		Input: userID is from user model
		Output: cart Model or error

		args: Find user cart or create new cart if it doesn't exist
	*/
	var cart cartModel.Cart

	// Try to find an existing cart
	err := r.db.Where("user_id = ?", userID).First(&cart).Error
	if err == nil {
		// Cart already exists, return it
		return &cart, nil
	}

	// If error is not "record not found", return the error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	// If no existing cart, create a new one
	newCart := cartModel.Cart{UserID: userID}
	if err := r.db.Create(&newCart).Error; err != nil {
		return nil, err
	}

	return &newCart, nil
}
