package service

import (
	"oolio/api-ecommerce/src/internal/model"
	"oolio/api-ecommerce/src/internal/repository"
)

type OrderService interface {
	PlaceOrder(items []model.OrderItem) (model.Order, error)
}

type OrderServiceImpl struct {
	Repo repository.OrderRepository
}

func NewOrderService(repo repository.OrderRepository) *OrderServiceImpl {
	return &OrderServiceImpl{Repo: repo}
}

func (s *OrderServiceImpl) PlaceOrder(items []model.OrderItem) (model.Order, error) {
	newOrder := model.Order{Items: items}

	createdOrder := s.Repo.CreateOrder(newOrder)

	return createdOrder, nil
}
