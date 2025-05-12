package repository

import (
	"golang-shop/internal/cart/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	// Migrate the schema
	err = db.AutoMigrate(&model.Cart{}, &model.CartItem{})
	assert.NoError(t, err)

	return db
}

func TestNewCartRepository(t *testing.T) {
	db := setupTestDB(t)
	repo := NewCartRepository(db)
	assert.NotNil(t, repo)
	assert.Equal(t, db, repo.db)
}

func TestFindOrCreateCart(t *testing.T) {
	db := setupTestDB(t)
	repo := NewCartRepository(db)

	t.Run("should create new cart when none exists", func(t *testing.T) {
		userID := uint(1)
		cart, err := repo.FindOrCreateCart(userID)

		assert.NoError(t, err)
		assert.NotNil(t, cart)
		assert.Equal(t, userID, cart.UserID)
	})

	t.Run("should return existing cart", func(t *testing.T) {
		userID := uint(2)

		// Create initial cart
		cart1, err := repo.FindOrCreateCart(userID)
		assert.NoError(t, err)

		// Try to find/create again
		cart2, err := repo.FindOrCreateCart(userID)
		assert.NoError(t, err)

		assert.Equal(t, cart1.ID, cart2.ID)
	})
}

func TestAddToCart(t *testing.T) {
	db := setupTestDB(t)
	repo := NewCartRepository(db)

	t.Run("should add item to cart", func(t *testing.T) {
		// Create a cart first
		cart, err := repo.FindOrCreateCart(1)
		assert.NoError(t, err)

		item := model.CartItem{
			ProductID: 1,
			Quantity:  2,
			Price:     10.99,
		}

		err = repo.AddToCart(cart.ID, item)
		assert.NoError(t, err)

		// Verify item was added
		retrievedCart, err := repo.GetCartItem(1)
		assert.NoError(t, err)
		assert.Len(t, retrievedCart.Items, 1)
		assert.Equal(t, item.ProductID, retrievedCart.Items[0].ProductID)
		assert.Equal(t, item.Quantity, retrievedCart.Items[0].Quantity)
	})
}

func TestUpdateItemQuantity(t *testing.T) {
	db := setupTestDB(t)
	repo := NewCartRepository(db)

	t.Run("should update item quantity", func(t *testing.T) {
		// Create cart and add item
		cart, err := repo.FindOrCreateCart(1)
		assert.NoError(t, err)

		item := model.CartItem{
			ProductID: 1,
			Quantity:  2,
			Price:     10.99,
		}
		err = repo.AddToCart(cart.ID, item)
		assert.NoError(t, err)

		// Update quantity
		newQuantity := 5
		err = repo.UpdateItemQuantity(1, cart.ID, newQuantity)
		assert.NoError(t, err)

		// Verify update
		retrievedCart, err := repo.GetCartItem(1)
		assert.NoError(t, err)
		assert.Len(t, retrievedCart.Items, 1)
		assert.Equal(t, newQuantity, retrievedCart.Items[0].Quantity)
	})
}

func TestRemoveCartItem(t *testing.T) {
	db := setupTestDB(t)
	repo := NewCartRepository(db)

	t.Run("should remove item from cart", func(t *testing.T) {
		// Create cart and add item
		cart, err := repo.FindOrCreateCart(1)
		assert.NoError(t, err)

		item := model.CartItem{
			ProductID: 1,
			Quantity:  2,
			Price:     10.99,
		}
		err = repo.AddToCart(cart.ID, item)
		assert.NoError(t, err)

		// Remove item
		err = repo.RemoveCartItem(cart.ID, 1)
		assert.NoError(t, err)

		// Verify removal
		retrievedCart, err := repo.GetCartItem(1)
		assert.NoError(t, err)
		assert.Len(t, retrievedCart.Items, 0)
	})
}

func TestClearCart(t *testing.T) {
	db := setupTestDB(t)
	repo := NewCartRepository(db)

	t.Run("should clear all items from cart", func(t *testing.T) {
		// Create cart and add multiple items
		cart, err := repo.FindOrCreateCart(1)
		assert.NoError(t, err)

		items := []model.CartItem{
			{ProductID: 1, Quantity: 2, Price: 10.99},
			{ProductID: 2, Quantity: 1, Price: 15.99},
		}

		for _, item := range items {
			err = repo.AddToCart(cart.ID, item)
			assert.NoError(t, err)
		}

		// Clear cart
		err = repo.ClearCart(cart.ID)
		assert.NoError(t, err)

		// Verify cart is empty
		retrievedCart, err := repo.GetCartItem(1)
		assert.NoError(t, err)
		assert.Len(t, retrievedCart.Items, 0)
	})
}
