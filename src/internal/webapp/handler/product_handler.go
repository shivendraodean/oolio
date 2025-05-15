package handler

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	"oolio/api-ecommerce/src/internal/service"
)

type ProductHandler struct {
	ProductService service.ProductService
}

func NewProductHandler(productService service.ProductService) *ProductHandler {
	return &ProductHandler{ProductService: productService}
}

func (h *ProductHandler) GetProductByID(c *gin.Context) {
	productID := c.Param("productId")

	if !(ValidateGetProductRequest(productID)) {
		slog.Error("GetProductRequest:Bad request: invalid product ID", "productID", productID)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	product := h.ProductService.GetProduct(productID)

	c.JSON(http.StatusOK, product)
}
