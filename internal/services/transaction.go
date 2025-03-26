package services

import "golang-shop/internal/repository"

type TransactionServices struct {
	cartRepo    *repository.CartRepository
	productRepo *repository.ProductRepository
}
