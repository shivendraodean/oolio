package service

import (
	"errors"
	"oolio/api-ecommerce/src/internal/model"
	"oolio/api-ecommerce/src/internal/repository"
)

type OrderService interface {
	PlaceOrder(items []model.OrderItem) (model.Order, error)
}

type OrderServiceImpl struct {
	Repo        repository.OrderRepository
	ProductRepo repository.ProductRepository
}

func NewOrderService(repo repository.OrderRepository, productRepo repository.ProductRepository) *OrderServiceImpl {
	return &OrderServiceImpl{Repo: repo, ProductRepo: productRepo}
}

func (s *OrderServiceImpl) PlaceOrder(items []model.OrderItem) (model.Order, error) {
	var products []model.Product

	for _, item := range items {
		product := s.ProductRepo.FindProduct(item.ProductID)
		if product.ID == "" {
			return model.Order{}, errors.New(ErrOrderedProductNotFoundMsg)
		}
		products = append(products, product)
	}

	newOrder := model.Order{Items: items, Products: products}

	createdOrder := s.Repo.CreateOrder(newOrder)

	return createdOrder, nil
}
