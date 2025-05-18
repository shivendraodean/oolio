package service

import (
	"errors"
	"oolio/api-ecommerce/src/internal/model"
	"oolio/api-ecommerce/src/internal/repository"
)

type ProductService interface {
	GetProduct(productID int64) (model.Product, error)
	ListProducts() ([]model.Product, error)
}

type ProductServiceImpl struct {
	Repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) *ProductServiceImpl {
	return &ProductServiceImpl{Repo: repo}
}

func (s *ProductServiceImpl) GetProduct(productID int64) (model.Product, error) {
	product := s.Repo.FindProduct(productID)
	if product.ID == "" {
		return model.Product{}, errors.New(ErrProductNotFoundMsg)
	}

	return product, nil
}

func (s *ProductServiceImpl) ListProducts() ([]model.Product, error) {
	products := s.Repo.ListProducts()
	return products, nil
}
