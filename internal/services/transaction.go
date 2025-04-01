package services

import "golang-shop/internal/repository"

type TransactionServices struct {
	TransactionRepo *repository.TransactionRepository
	cartRepo    *repository.CartRepository
	productRepo *repository.ProductRepository
}

func NewTransactionRepository(
	trasactionRepo *repository.TransactionRepository,
	cartRepo *repository.CartRepository,
	productRepo *repository.ProductRepository,
) *TransactionServices {
	return &TransactionServices{TransactionRepo: trasactionRepo,cartRepo: cartRepo, productRepo: productRepo}
}


