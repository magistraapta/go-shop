package services

import (
	"errors"
	"golang-shop/internal/dto"
	"golang-shop/internal/model"
	"golang-shop/internal/repository"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(userRequest model.User) error {
	return s.repo.CreateUser(userRequest)
}

func (s *UserService) Login(userLogin dto.UserLogin) (string, error) {
	// get user
	user, err := s.repo.Login(userLogin)
	if err != nil {
		return "", errors.New("Invalid email")
	}

	// hash user password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userLogin.Password)); err != nil {
		return "", errors.New("Failed to compare password")
	}
	// sign jwt

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(10 * time.Minute).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		return "", errors.New("Failed to create JWT Token")
	}

	return tokenString, nil
}

func (s *UserService) GetUserById(id int) (*dto.UserResponse, error) {
	return s.repo.GetUserById(id)
}
