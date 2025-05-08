package repository

import (
	"golang-shop/internal/auth/dto"
	"golang-shop/internal/auth/model"

	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) CreateUser(userRequest model.User) error {

	if err := r.db.Create(&userRequest).Error; err != nil {
		return err
	}

	return nil
}

func (r *AuthRepository) CheckUserExist(loginRequest dto.LoginRequest) (*model.User, error) {
	var user model.User

	if err := r.db.Where("email = ?", loginRequest.Email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *AuthRepository) GetUserById(id int) (*model.User, error) {
	var user model.User

	result := r.db.Find(&user, id)

	if result.Error != nil {
		return nil, result.Error
	}

	userResponse := model.User{
		Username: user.Username,
		Email:    user.Email,
	}

	return &userResponse, nil
}
