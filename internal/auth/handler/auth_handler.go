package handler

import (
	"golang-shop/internal/auth/dto"
	"golang-shop/internal/auth/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service *services.AuthService
}

func NewAuthHandler(service *services.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

// @Summary Register User
// @Description Register a new user
// @Tags auth
// @Accept json
// @Produce json
// @Param username body dto.RegisterRequest true "Register Request"
// @Success 201 {object} dto.AuthResponse
// @Failure 400 {object} dto.AuthResponse
// @Failure 500 {object} dto.AuthResponse
// @Router /user/register [post]
func (h *AuthHandler) RegisterUser(ctx *gin.Context) {
	var request dto.RegisterRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.AuthResponse{Message: "Invalid body request"})
		return
	}

	if err := h.service.Register(request.Username, request.Email, request.Password, "user"); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, dto.AuthResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, dto.AuthResponse{Message: "Register Success"})
}

// @Summary Register Admin
// @Description Register a new admin
// @Tags auth
// @Accept json
// @Produce json
// @Param username body dto.RegisterRequest true "Register Request"
// @Success 201 {object} dto.AuthResponse
// @Failure 400 {object} dto.AuthResponse
// @Failure 500 {object} dto.AuthResponse
// @Router /admin/register [post]
func (h *AuthHandler) RegisterAdmin(ctx *gin.Context) {
	var request dto.RegisterRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.AuthResponse{Message: "Invalid body request"})
		return
	}

	if err := h.service.Register(request.Username, request.Email, request.Password, "admin"); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, dto.AuthResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, dto.AuthResponse{Message: "Register Success"})
}

// @Summary Login
// @Description Login to the system
// @Tags auth
// @Accept json
// @Produce json
// @Param email body dto.LoginRequest true "Login Request"
// @Success 200 {object} dto.AuthResponse
// @Failure 400 {object} dto.AuthResponse
// @Failure 500 {object} dto.AuthResponse
// @Router /auth/login [post]
func (h *AuthHandler) Login(ctx *gin.Context) {
	var request dto.LoginRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.AuthResponse{Message: "Invalid body request"})
		return
	}

	token, err := h.service.Login(request)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, dto.AuthResponse{Message: err.Error()})
		return
	}
	ctx.SetCookie("Authorization", token, 3600, "/", "localhost", false, true)
	ctx.JSON(http.StatusOK, dto.AuthResponse{Message: "Login Success"})

}
