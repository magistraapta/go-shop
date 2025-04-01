package test

import (
	"golang-shop/internal/model"
	"golang-shop/internal/repository"
	"golang-shop/internal/services"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&model.User{})

	repo := repository.NewUserRepository(db)
	services := services.NewUserService(repo)

	mockCreateUser := model.User{
		Username: "mikeee",
		Email:    "mike@email.com",
		Password: "mike123",
	}

	t.Run("Signup should not return an error", func(t *testing.T) {
		err := services.CreateUser(mockCreateUser)
		assert.Nil(t, err, "Signup should not return an error")
	})

	t.Run("User should exist in database", func(t *testing.T) {
		var createdUser model.User
		result := db.First(&createdUser, "username = ?", mockCreateUser.Username)
		assert.Nil(t, result.Error, "User should exist in the database")

		// Ensure stored user has expected values
		assert.Equal(t, mockCreateUser.Username, createdUser.Username)

	})
}
