package repository

import (
	"testing"

	"golang-shop/internal/auth/dto"
	"golang-shop/internal/auth/model"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	// Create in-memory SQLite database
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to open sqlite db: %v", err)
	}

	// Auto migrate the schema
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		t.Fatalf("Failed to migrate schema: %v", err)
	}

	return db
}

func TestAuthRepository_CreateUser(t *testing.T) {
	db := setupTestDB(t)
	repo := NewAuthRepository(db)

	user := model.User{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "hashedpassword",
		Role:     "user",
	}

	// Test the function
	err := repo.CreateUser(user)
	assert.NoError(t, err)

	// Verify the user was created
	var createdUser model.User
	result := db.First(&createdUser, "email = ?", user.Email)
	assert.NoError(t, result.Error)
	assert.Equal(t, user.Username, createdUser.Username)
	assert.Equal(t, user.Email, createdUser.Email)
	assert.Equal(t, user.Password, createdUser.Password)
	assert.Equal(t, user.Role, createdUser.Role)
}

func TestAuthRepository_CheckUserExist(t *testing.T) {
	db := setupTestDB(t)
	repo := NewAuthRepository(db)

	// Create a test user first
	testUser := model.User{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "hashedpassword",
		Role:     "user",
	}
	err := db.Create(&testUser).Error
	assert.NoError(t, err)

	loginRequest := dto.LoginRequest{
		Email:    "test@example.com",
		Password: "password123",
	}

	// Test the function
	user, err := repo.CheckUserExist(loginRequest)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, testUser.Email, user.Email)
	assert.Equal(t, testUser.Username, user.Username)
}

func TestAuthRepository_GetUserById(t *testing.T) {
	db := setupTestDB(t)
	repo := NewAuthRepository(db)

	// Create a test user first
	testUser := model.User{
		Username: "testuser",
		Email:    "test@example.com",
		Role:     "user",
	}
	err := db.Create(&testUser).Error
	assert.NoError(t, err)

	// Test the function
	user, err := repo.GetUserById(int(testUser.ID))
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, testUser.Email, user.Email)
	assert.Equal(t, testUser.Username, user.Username)
}
