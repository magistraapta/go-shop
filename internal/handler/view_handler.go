package handler

import (
	"golang-shop/internal/services"
	"golang-shop/internal/views"
	"net/http"
	"strconv"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

type HtmlHandler struct {
	userService     *services.UserService
	productServices *services.ProductServices
}

func NewHtmlHandler(userService *services.UserService, productServices *services.ProductServices) *HtmlHandler {
	return &HtmlHandler{userService: userService, productServices: productServices}
}

func (h *HtmlHandler) RenderHome(c *gin.Context) {
	products, err := h.productServices.GetAllProduct()

	if err != nil {
		c.Status(http.StatusInternalServerError)
		templ.Handler(views.ErrorPage(err)).ServeHTTP(c.Writer, c.Request)
		return
	}

	templ.Handler(views.Home(*products)).ServeHTTP(c.Writer, c.Request)
}

func (h *HtmlHandler) ProductDetail(c *gin.Context) {
	productID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Status(http.StatusBadRequest)
		templ.Handler(views.ErrorPage(err))
		return
	}
	product, err := h.productServices.GetProductById(productID)

	if err != nil {
		c.Status(http.StatusInternalServerError)
		templ.Handler(views.ErrorPage(err))
		return
	}

	templ.Handler(views.ProductDetail(*product)).ServeHTTP(c.Writer, c.Request)
}
