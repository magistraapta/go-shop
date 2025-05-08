package repository

import (
	"golang-shop/internal/product/model"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) CreateProduct(productRequest model.Product) error {
	result := r.db.Create(&productRequest)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *ProductRepository) DeleteProductById(id int) error {
	var product model.Product

	if err := r.db.Delete(&product, id).Error; err != nil {
		return err
	}

	return nil
}

func (r *ProductRepository) GetProductById(id int) (*model.Product, error) {
	var product model.Product

	result := r.db.Find(&product, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &product, nil
}

func (r *ProductRepository) GetAllProduct() (*[]model.Product, error) {
	var products []model.Product

	result := r.db.Find(&products)

	if result.Error != nil {
		return nil, result.Error
	}

	return &products, nil
}

func (r *ProductRepository) UpdateProduct(productID uint, updateRequest model.Product) (*model.Product, error) {
	// Find the product by ID
	var product model.Product
	if err := r.db.First(&product, productID).Error; err != nil {
		return nil, err
	}

	// Apply updates only to provided fields
	if err := r.db.Model(&product).Updates(updateRequest).Error; err != nil {
		return nil, err
	}

	// Return the actual updated product
	return &product, nil
}

func (r *ProductRepository) UpdateProductStock(productID uint, quantity int) error {
	if err := r.db.Model(&model.Product{}).
		Where("id = ?", productID).
		Update("stock", gorm.Expr("stock - ?", quantity)).Error; err != nil {
		return err
	}

	return nil
}
