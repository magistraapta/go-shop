package repository

import (
	"golang-shop/internal/model"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) CreateTransaction(transaction model.Transaction, items []model.OrderItem, paymentMethod string) error {
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
