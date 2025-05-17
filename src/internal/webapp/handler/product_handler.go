package handler

import (
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
	productID, err := ValidateGetProductRequest(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID supplied"})

		return
	}

	product, err := h.ProductService.GetProduct(productID)

	if err != nil {
		if err.Error() == service.ErrProductNotFoundMsg {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "An error occurred"})
		}
		return
	}

	c.JSON(http.StatusOK, product)
}
