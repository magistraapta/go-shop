package handler

import (
	"golang-shop/internal/auth/model"
	"golang-shop/internal/cart/dto"
	"golang-shop/internal/cart/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	services *services.CartServices
}

func NewCartHandler(services *services.CartServices) *CartHandler {
	return &CartHandler{services: services}
}

func (h *CartHandler) AddToCart(c *gin.Context) {
	var request dto.AddToCartRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	// get authenticated user
	user, exist := c.Get("user")

	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	currentUser, ok := user.(model.User)

	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Invalid user",
		})
	}

	err := h.services.AddToCart(currentUser.ID, request)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to add item to cart",
			"error":   err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item added to cart successfully"})
}

func (h *CartHandler) GetCart(c *gin.Context) {

	user, exist := c.Get("user")

	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	currentUser, ok := user.(model.User)

	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Invalid user",
		})
	}

	cart, err := h.services.GetCartItems(currentUser.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get cart",
			"error":   err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"cart": cart,
	})
}

func (h *CartHandler) UpdateQuantity(c *gin.Context) {

	user, exist := c.Get("user")

	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	currentUser, ok := user.(model.User)

	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Invalid user",
		})
	}

	productID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	var request struct {
		Quantity int `json:"quantity" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.services.UpdateQuantity(currentUser.ID, uint(productID), request.Quantity); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cart item quantity updated"})
}

func (h *CartHandler) RemoveItem(c *gin.Context) {

	user, exist := c.Get("user")

	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	currentUser, ok := user.(model.User)

	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Invalid user",
		})
	}

	itemID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
		return
	}

	if err := h.services.RemoveItem(currentUser.ID, uint(itemID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item removed from cart"})
}
