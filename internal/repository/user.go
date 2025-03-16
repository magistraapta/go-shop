package repository

import (
	"golang-shop/internal/dto"
	"golang-shop/internal/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(userRequest model.User) error {
	result := r.db.Create(&userRequest)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *UserRepository) Login(loginRequest dto.UserLogin) (*model.User, error) {
	var user model.User

	if err := r.db.Where("email = ?", loginRequest.Email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
