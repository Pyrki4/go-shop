package handlers

import (
	"github.com/Pyrki4/go-shop/internal/services"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productService *services.ProductService
}

func NewProductHandler(productService *services.ProductService) *ProductHandler {
	return &ProductHandler{productService: productService}
}

func (h *ProductHandler) GetAllProducts(c *gin.Context) {
	products := h.productService.GetAllProducts()
	c.JSON(200, products)
}

func RegisterRoutes(r *gin.Engine, productService *services.ProductService) {
	handler := NewProductHandler(productService)
	r.GET("/products", handler.GetAllProducts)
}
