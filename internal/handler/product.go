package handler

import (
	"golang-shop/internal/model"
	"golang-shop/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	services *services.ProductServices
}

func NewProductHandler(services *services.ProductServices) *ProductHandler {
	return &ProductHandler{services: services}
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var productRequest model.Product

	if err := c.ShouldBindBodyWithJSON(&productRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid body request",
		})
	}

	err := h.services.CreateProduct(productRequest)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to create product",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success create new product",
	})

}

func (h *ProductHandler) DeleteProductById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to get id",
		})
	}

	if err := h.services.DeleteProductById(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to delete product",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success delete product",
	})
}

func (h *ProductHandler) GetProductById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to get id",
		})
	}

	product, err := h.services.GetProductById(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get product",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success getting product",
		"product": product,
	})
}

func (h *ProductHandler) GetAllProduct(c *gin.Context) {
	products, err := h.services.GetAllProduct()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get all product",
			"error":   err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Successfully getting all product",
		"products": products,
	})
}

func (h *ProductHandler) UpdateProductById(c *gin.Context) {
	// get user by id
	productID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid product id",
			"error":   err.Error(),
		})
		return
	}

	var updateRequest model.Product

	if err := c.ShouldBindJSON(&updateRequest); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Invalid Request body",
		})
	}

	updatedProduct, err := h.services.UpdateProductById(uint(productID), updateRequest)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully update product",
		"product": updatedProduct,
	})
}
