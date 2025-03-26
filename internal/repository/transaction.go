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

func (r *TransactionRepository) CreateTransaction(transaction model.Transaction, items []model.OrderItem) error {
	if err := r.db.Create(&transaction).Error; err != nil {
		return err

	}

	for i := range items {
		items[i].ID = transaction.ID
		if err := r.db.Create(&items[i]).Error; err != nil {
			return err
		}
	}

	return nil
}
