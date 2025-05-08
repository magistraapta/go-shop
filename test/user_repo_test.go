package test

import (
	"testing"

	"golang-shop/internal/auth/dto"
	"golang-shop/internal/auth/model"
	"golang-shop/internal/auth/repository"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	// Migrate the schema
	err = db.AutoMigrate(&model.User{})
	assert.NoError(t, err)

	return db
}

func TestCreateUser(t *testing.T) {
	db := setupTestDB(t)
	repo := repository.NewAuthRepository(db)

	user := model.User{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password123",
	}

	err := repo.CreateUser(user)
	assert.NoError(t, err)

	// Verify user was created
	var createdUser model.User
	err = db.Where("email = ?", user.Email).First(&createdUser).Error
	assert.NoError(t, err)
	assert.Equal(t, user.Username, createdUser.Username)
	assert.Equal(t, user.Email, createdUser.Email)
}

func TestCheckUserExist(t *testing.T) {
	db := setupTestDB(t)
	repo := repository.NewAuthRepository(db)

	// Create a test user
	user := model.User{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password123",
	}
	err := db.Create(&user).Error
	assert.NoError(t, err)

	// Test existing user
	loginRequest := dto.LoginRequest{
		Email: "test@example.com",
	}
	foundUser, err := repo.CheckUserExist(loginRequest)
	assert.NoError(t, err)
	assert.NotNil(t, foundUser)
	assert.Equal(t, user.Email, foundUser.Email)

	// Test non-existing user
	nonExistingRequest := dto.LoginRequest{
		Email: "nonexisting@example.com",
	}
	_, err = repo.CheckUserExist(nonExistingRequest)
	assert.Error(t, err)
}

func TestGetUserById(t *testing.T) {
	db := setupTestDB(t)
	repo := repository.NewAuthRepository(db)

	// Create a test user
	user := model.User{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password123",
	}
	err := db.Create(&user).Error
	assert.NoError(t, err)

	// Test getting existing user
	foundUser, err := repo.GetUserById(int(user.ID))
	assert.NoError(t, err)
	assert.NotNil(t, foundUser)
	assert.Equal(t, user.Username, foundUser.Username)
	assert.Equal(t, user.Email, foundUser.Email)

	// Test getting non-existing user
	nonExistingUser, err := repo.GetUserById(999)
	assert.NoError(t, err)
	assert.NotNil(t, nonExistingUser)
	assert.Empty(t, nonExistingUser.Username)
	assert.Empty(t, nonExistingUser.Email)
}
