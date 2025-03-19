package handler

import (
	"golang-shop/internal/dto"
	"golang-shop/internal/model"
	"golang-shop/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	services *services.UserService
}

func NewUserHandler(services *services.UserService) *UserHandler {
	return &UserHandler{services: services}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var userRequest model.User

	if err := c.ShouldBindBodyWithJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Error getting request body",
		})
	}

	// hash user password
	hash, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to hash password",
		})
		return
	}

	userRequest.Password = string(hash)

	if err := h.services.CreateUser(userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Failed to create user",
			"err": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User successfully created!",
	})
}

func (h *UserHandler) TestApi(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "hello world",
	})
}

func (h *UserHandler) Login(c *gin.Context) {
	var userLogin dto.UserLogin

	if err := c.ShouldBindBodyWithJSON(&userLogin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Invalid request body",
		})

		return
	}

	token, err := h.services.Login(userLogin)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", token, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"accessToken": token,
	})
}

func (h *UserHandler) GetUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Failed to get id",
		})
	}

	user, err := h.services.GetUserById(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Failed to get user",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "Success",
		"user": user,
	})
}
