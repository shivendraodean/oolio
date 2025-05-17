package webapp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter(container *Container) *gin.Engine {
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hey, I am Healthy!",
		})
	})

	// Product routes
	router.GET("/product/:productId", container.ProductHandler.GetProductByID)

	// Order routes
	router.POST("/order", container.OrderHandler.PlaceOrder)

	return router
}
