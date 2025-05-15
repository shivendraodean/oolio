package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"oolio/api-ecommerce/src/internal/handler"
	"oolio/api-ecommerce/src/internal/repository"
	"oolio/api-ecommerce/src/internal/service"
)

func main() {
	router := gin.Default()

	// Initialize dependencies
	productRepo := repository.NewProductRepository()
	productService := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hey, I am Healthy!",
		})
	})

	// Product routes
	router.GET("/product/:productId", productHandler.GetProductByID)

	router.Run(":8080")
}
