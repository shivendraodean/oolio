package service

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"oolio/api-ecommerce/src/internal/model"
)

type OrderRepoMock struct {
	CreateOrderFunc func(order model.Order) model.Order
}

type ProductRepoMock struct {
	FindProductFunc func(id int64) model.Product
}

func (m *OrderRepoMock) CreateOrder(order model.Order) model.Order {
	return m.CreateOrderFunc(order)
}

func (m *ProductRepoMock) FindProduct(id int64) model.Product {
	return m.FindProductFunc(id)
}

func TestOrderService_PlaceOrder(t *testing.T) {
	t.Run("should successfully place an order", func(t *testing.T) {
		orderRepo := &OrderRepoMock{}
		productRepo := &ProductRepoMock{}
		orderRepo.CreateOrderFunc = func(order model.Order) model.Order {
			order.ID = "123"
			return order
		}
		productRepo.FindProductFunc = func(id int64) model.Product {
			return model.Product{
				ID: "1",
			}
		}
		orderItems := []model.OrderItem{{ProductID: 10, Quantity: 1}}
		service := NewOrderService(orderRepo, productRepo)

		order, err := service.PlaceOrder(orderItems)

		assert.NoError(t, err)
		assert.NotNil(t, order)
		assert.Equal(t, "123", order.ID)
		assert.Equal(t, 1, len(order.Items))
	})

	t.Run("should return error when product is not found and do not place order", func(t *testing.T) {
		orderRepo := &OrderRepoMock{}
		productRepo := &ProductRepoMock{}

		productRepo.FindProductFunc = func(id int64) model.Product {
			return model.Product{}
		}
		orderRepo.CreateOrderFunc = func(order model.Order) model.Order {
			return model.Order{}
		}

		orderItems := []model.OrderItem{{ProductID: -1, Quantity: 1}}
		service := NewOrderService(orderRepo, productRepo)

		_, err := service.PlaceOrder(orderItems)

		assert.Error(t, err)
		assert.Equal(t, ErrOrderedProductNotFoundMsg, err.Error())
	})
}
