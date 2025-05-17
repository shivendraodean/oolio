package handler

import (
	"net/http"
	"oolio/api-ecommerce/src/internal/model"
	"oolio/api-ecommerce/src/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	OrderService service.OrderService
}

func NewOrderHandler(orderService service.OrderService) *OrderHandler {
	return &OrderHandler{OrderService: orderService}
}

func (h *OrderHandler) PlaceOrder(c *gin.Context) {
	orderRequest, err := ValidatePlaceOrderRequest(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order request"})

		return
	}

	var items []model.OrderItem
	for _, item := range orderRequest.Items {
		productID, _ := strconv.ParseInt(item.ProductID, 10, 64)
		items = append(items, model.OrderItem{ProductID: productID, Quantity: item.Quantity})
	}

	order, err := h.OrderService.PlaceOrder(items, orderRequest.CouponCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"An error occurred": err.Error()})

		return
	}

	c.JSON(http.StatusCreated, order)
}
