package service

import (
	"oolio/api-ecommerce/src/internal/model"
	"oolio/api-ecommerce/src/internal/repository"
)

type ProductService interface {
	GetProduct(productID string) model.Product
}

type ProductServiceImpl struct {
	Repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) *ProductServiceImpl {
	return &ProductServiceImpl{Repo: repo}
}

func (s *ProductServiceImpl) GetProduct(productID string) model.Product {
	return s.Repo.FindProduct(productID)
}
