package service

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"oolio/api-ecommerce/src/internal/model"
)

type MockProductRepository struct {
	FindProductFunc func(id int64) model.Product
}

func (m *MockProductRepository) FindProduct(id int64) model.Product {
	return m.FindProductFunc(id)
}

func TestProductService_GetProduct(t *testing.T) {
	t.Run("should return the correct product", func(t *testing.T) {
		expectedProduct := model.Product{
			ID:       "123",
			Name:     "Test Product",
			Price:    9.99,
			Category: "Test Category",
			Image: model.ProductImage{
				Thumbnail: "thumbnail.jpg",
				Mobile:    "mobile.jpg",
				Tablet:    "tablet.jpg",
				Desktop:   "desktop.jpg",
			},
		}

		mockRepo := &MockProductRepository{
			FindProductFunc: func(id int64) model.Product {
				return expectedProduct
			},
		}

		service := NewProductService(mockRepo)

		result, err := service.GetProduct(123)

		assert.NoError(t, err)
		assert.Equal(t, expectedProduct, result)
	})

	t.Run("should return not found error when product data is nil", func(t *testing.T) {
		mockRepo := &MockProductRepository{
			FindProductFunc: func(id int64) model.Product {
				return model.Product{}
			},
		}

		service := NewProductService(mockRepo)

		_, err := service.GetProduct(0)

		assert.Error(t, err)
		assert.Equal(t, ErrProductNotFoundMsg, err.Error())
	})
}
