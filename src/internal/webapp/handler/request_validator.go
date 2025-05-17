package handler

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

// Products

type ProductRequest struct {
	ID int64 `uri:"productId" binding:"required,min=1"`
}

func ValidateGetProductRequest(c *gin.Context) (int64, error) {
	var request ProductRequest

	if err := c.ShouldBindUri(&request); err != nil {
		slog.Error("Invalid product ID supplied", "error", err.Error())
		return 0, err
	}

	return request.ID, nil
}

// Orders

type OrderItemRequest struct {
	ProductID string `json:"productId" binding:"required"`
	Quantity  int    `json:"quantity" binding:"required"`
}

type PlaceOrderRequest struct {
	CouponCode string             `json:"couponCode" binding:"omitempty,min=8,max=10"`
	Items      []OrderItemRequest `json:"items" binding:"required,dive"`
}

func ValidatePlaceOrderRequest(c *gin.Context) (PlaceOrderRequest, error) {
	var request PlaceOrderRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		slog.Error("Invalid place order request", "error", err.Error())
		return PlaceOrderRequest{}, err
	}

	return request, nil
}
