package services

import (
	"golang-shop/internal/product/dto"
	"golang-shop/internal/product/model"
	"golang-shop/internal/product/repository"
)

type ProductServices struct {
	repo *repository.ProductRepository
}

func NewProductServices(repo *repository.ProductRepository) *ProductServices {
	return &ProductServices{repo: repo}
}

func (s *ProductServices) CreateProduct(productRequest dto.ProductRequest) error {
	return s.repo.CreateProduct(productRequest)
}

func (s *ProductServices) DeleteProductById(id int) error {
	return s.repo.DeleteProductById(id)
}

func (s *ProductServices) GetProductById(id int) (*model.Product, error) {
	return s.repo.GetProductById(id)
}

func (s *ProductServices) GetAllProduct() (*[]model.Product, error) {
	return s.repo.GetAllProduct()
}

func (s *ProductServices) UpdateProductById(productID uint, updateProductRequest model.Product) (*model.Product, error) {
	return s.repo.UpdateProduct(productID, updateProductRequest)
}
