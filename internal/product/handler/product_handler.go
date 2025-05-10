package handler

import (
	"golang-shop/internal/product/dto"
	"golang-shop/internal/product/model"
	"golang-shop/internal/product/services"
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

// @Summary Create Product
// @Description Create a new product
// @Tags products
// @Accept json
// @Produce json
// @Param product body dto.ProductRequest true "Product"
// @Success 201 {object} dto.Message
// @Failure 400 {object} dto.Message
// @Failure 500 {object} dto.Message
// @Router /product [post]
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var productRequest dto.ProductRequest

	if err := c.ShouldBindBodyWithJSON(&productRequest); err != nil {
		c.JSON(http.StatusBadRequest, dto.Message{
			Message: "Invalid body request",
		})
	}

	err := h.services.CreateProduct(productRequest)

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Message{
			Message: "failed to create product",
		})
		return
	}

	c.JSON(http.StatusOK, dto.Message{
		Message: "Success create new product",
	})
}

// @Summary Delete Product
// @Description Delete a product by id
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} dto.Message
// @Failure 400 {object} dto.Message
// @Failure 500 {object} dto.Message
// @Router /product/{id} [delete]
func (h *ProductHandler) DeleteProductById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Message{
			Message: "Failed to get id",
		})
	}

	if err := h.services.DeleteProductById(id); err != nil {
		c.JSON(http.StatusInternalServerError, dto.Message{
			Message: "Failed to delete product",
		})
	}

	c.JSON(http.StatusOK, dto.Message{
		Message: "Success delete product",
	})
}

// @Summary Get Product By Id
// @Description Get a product by id
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} dto.ProductResponse
// @Failure 400 {object} dto.Message
// @Failure 500 {object} dto.Message
// @Router /product/{id} [get]

func (h *ProductHandler) GetProductById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Message{
			Message: "Failed to get id",
		})
	}

	product, err := h.services.GetProductById(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Message{
			Message: "Failed to get product",
		})
	}

	c.JSON(http.StatusOK, dto.ProductResponse{
		Message: "Success getting product",
		Data:    *product,
	})
}

// @Summary Get All Product
// @Description Get all products
// @Tags products
// @Accept json
// @Produce json
// @Success 200 {object} dto.ProductResponseList
// @Failure 500 {object} dto.Message
// @Router /product [get]
func (h *ProductHandler) GetAllProduct(c *gin.Context) {
	products, err := h.services.GetAllProduct()

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Message{
			Message: "Failed to get all product",
		})
	}

	c.JSON(http.StatusOK, dto.ProductResponseList{
		Message: "Successfully getting all product",
		Data:    *products,
	})
}

// @Summary Update Product By Id
// @Description Update a product by id
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} dto.ProductResponse
// @Failure 400 {object} dto.Message
// @Failure 500 {object} dto.Message
// @Router /product/{id} [put]
func (h *ProductHandler) UpdateProductById(c *gin.Context) {
	// get user by id
	productID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Message{
			Message: "Invalid product id",
		})
		return
	}

	var updateRequest model.Product

	if err := c.ShouldBindJSON(&updateRequest); err != nil {
		c.JSON(http.StatusInternalServerError, dto.Message{
			Message: "Invalid Request body",
		})
	}

	updatedProduct, err := h.services.UpdateProductById(uint(productID), updateRequest)

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Message{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.ProductResponse{
		Message: "Successfully update product",
		Data:    *updatedProduct,
	})
}
