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

// @Summary Add to cart
// @Description Add a product to the cart
// @Accept json
// @Produce json
// @Tags cart
// @Param addToCartRequest body dto.AddToCartRequest true "Add to cart request"
// @Success 200 {object} dto.CartResponse
// @Failure 401 {object} dto.CartResponse
// @Failure 500 {object} dto.CartResponse
// @Router /cart/add [post]
func (h *CartHandler) AddToCart(c *gin.Context) {
	var request dto.AddToCartRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, dto.CartResponse{
			Message: "Invalid request body",
		})
		return
	}

	// get authenticated user
	user, exist := c.Get("user")

	if !exist {
		c.JSON(http.StatusUnauthorized, dto.CartResponse{
			Message: "Unauthorized",
		})
		return
	}

	currentUser, ok := user.(model.User)

	if !ok {
		c.JSON(http.StatusInternalServerError, dto.CartResponse{
			Message: "Invalid user",
		})
	}

	err := h.services.AddToCart(currentUser.ID, request)

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.CartResponse{
			Message: "Failed to add item to cart",
		})
		return
	}

	c.JSON(http.StatusOK, dto.CartResponse{
		Message: "Item added to cart successfully",
	})
}

// @Summary Get cart items
// @Description Get all items in the cart
// @Accept json
// @Produce json
// @Tags cart
// @Success 200 {object} dto.CartResponse
// @Failure 401 {object} dto.CartResponse
// @Failure 500 {object} dto.CartResponse
// @Router /cart [get]
func (h *CartHandler) GetCart(c *gin.Context) {

	user, exist := c.Get("user")

	if !exist {
		c.JSON(http.StatusUnauthorized, dto.CartResponse{
			Message: "Unauthorized",
		})
		return
	}

	currentUser, ok := user.(model.User)

	if !ok {
		c.JSON(http.StatusInternalServerError, dto.CartResponse{
			Message: "Invalid user",
		})
	}

	cart, err := h.services.GetCartItems(currentUser.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.CartResponse{
			Message: "Failed to get cart",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"cart": cart,
	})
}

// @Summary Update quantity
// @Description Update the quantity of a product in the cart
// @Accept json
// @Produce json
// @Tags cart
// @Param updateQuantityRequest body dto.UpdateQuantityRequest true "Update quantity request"
// @Success 200 {object} dto.CartResponse
// @Failure 401 {object} dto.CartResponse
// @Failure 500 {object} dto.CartResponse
// @Router /cart [put]
func (h *CartHandler) UpdateQuantity(c *gin.Context) {

	user, exist := c.Get("user")

	if !exist {
		c.JSON(http.StatusUnauthorized, dto.CartResponse{
			Message: "Unauthorized",
		})
		return
	}

	currentUser, ok := user.(model.User)

	if !ok {
		c.JSON(http.StatusInternalServerError, dto.CartResponse{
			Message: "Invalid user",
		})
	}

	productID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, dto.CartResponse{
			Message: "Invalid product ID",
		})
		return
	}

	var request dto.UpdateQuantityRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, dto.CartResponse{
			Message: err.Error(),
		})
		return
	}

	if err := h.services.UpdateQuantity(currentUser.ID, uint(productID), request.Quantity); err != nil {
		c.JSON(http.StatusInternalServerError, dto.CartResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.CartResponse{
		Message: "Cart item quantity updated",
	})
}

// @Summary Remove item
// @Description Remove an item from the cart
// @Accept json
// @Produce json
// @Tags cart
// @Param id path uint true "Product ID"
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /cart/{id} [delete]
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
