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

func (r *UserRepository) GetUserById(id int) (*dto.UserResponse, error) {
	var user model.User

	result := r.db.Find(&user, id)

	if result.Error != nil {
		return nil, result.Error
	}

	userResponse := dto.UserResponse{
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return &userResponse, nil
}
