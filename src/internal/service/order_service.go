package service

import (
	"errors"
	"oolio/api-ecommerce/src/internal/model"
	"oolio/api-ecommerce/src/internal/repository"
)

type OrderService interface {
	PlaceOrder(items []model.OrderItem, couponCode string) (model.Order, error)
}

type OrderServiceImpl struct {
	Repo        repository.OrderRepository
	ProductRepo repository.ProductRepository
	CouponRepo  repository.CouponRepository
}

func NewOrderService(repo repository.OrderRepository, productRepo repository.ProductRepository, couponRepo repository.CouponRepository) *OrderServiceImpl {
	return &OrderServiceImpl{Repo: repo, ProductRepo: productRepo, CouponRepo: couponRepo}
}

func (s *OrderServiceImpl) PlaceOrder(items []model.OrderItem, couponCode string) (model.Order, error) {
	var products []model.Product

	for _, item := range items {
		product := s.ProductRepo.FindProduct(item.ProductID)
		if product.ID == "" {
			return model.Order{}, errors.New(ErrOrderedProductNotFoundMsg)
		}
		products = append(products, product)
	}

	newOrder := model.Order{Items: items, Products: products}

	if couponCode != "" {
		couponOccurences := s.CouponRepo.SearchCoupon(couponCode)

		if couponOccurences == 2 {
			newOrder.CouponCode = couponCode
		} else {
			return model.Order{}, errors.New(ErrInvalidCouponMsg)
		}
	}

	createdOrder := s.Repo.CreateOrder(newOrder)

	return createdOrder, nil
}
