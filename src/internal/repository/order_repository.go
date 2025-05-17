package repository

import "oolio/api-ecommerce/src/internal/model"

type OrderRepository interface {
	CreateOrder(order model.Order) model.Order
}

type OrderRepositoryImpl struct{}

func NewOrderRepository() OrderRepository {
	return &OrderRepositoryImpl{}
}

func (r *OrderRepositoryImpl) CreateOrder(order model.Order) model.Order {
	return order
}
