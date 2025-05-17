package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"oolio/api-ecommerce/src/internal/model"
	"oolio/api-ecommerce/src/internal/service"
)

type MockProductService struct {
	GetProductFunc func(productID int64) (model.Product, error)
}

func (m *MockProductService) GetProduct(productID int64) (model.Product, error) {
	return m.GetProductFunc(productID)
}

func TestProductHandler_GetProductByID(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		productID      string
		mockProduct    model.Product
		mockError      error
		expectedStatus int
		expectedBody   map[string]interface{}
	}{
		{
			name:      "Valid product ID returns product",
			productID: "123",
			mockProduct: model.Product{
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
			},
			mockError:      nil,
			expectedStatus: http.StatusOK,
			expectedBody: map[string]interface{}{
				"id":       "123",
				"name":     "Test Product",
				"price":    9.99,
				"category": "Test Category",
				"image": map[string]interface{}{
					"thumbnail": "thumbnail.jpg",
					"mobile":    "mobile.jpg",
					"tablet":    "tablet.jpg",
					"desktop":   "desktop.jpg",
				},
			},
		},
		{
			name:           "Invalid product ID format returns 400",
			productID:      "abc",
			mockProduct:    model.Product{},
			mockError:      nil, // Won't be called due to validation error
			expectedStatus: http.StatusBadRequest,
			expectedBody: map[string]interface{}{
				"error": "Invalid ID supplied",
			},
		},
		{
			name:           "Product not found returns 404",
			productID:      "999",
			mockProduct:    model.Product{},
			mockError:      errors.New(service.ErrProductNotFoundMsg),
			expectedStatus: http.StatusNotFound,
			expectedBody: map[string]interface{}{
				"error": "Product not found",
			},
		},
		{
			name:           "Service error returns 500",
			productID:      "123",
			mockProduct:    model.Product{},
			mockError:      errors.New("database error"),
			expectedStatus: http.StatusInternalServerError,
			expectedBody: map[string]interface{}{
				"error": "An error occurred",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			mockService := &MockProductService{}
			mockService.GetProductFunc = func(productID int64) (model.Product, error) { return tt.mockProduct, tt.mockError }

			handler := NewProductHandler(mockService)

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = httptest.NewRequest(http.MethodGet, "/products/"+tt.productID, nil)
			c.Params = []gin.Param{{Key: "productId", Value: tt.productID}}

			// Act
			handler.GetProductByID(c)

			// Assert
			assert.Equal(t, tt.expectedStatus, w.Code)

			var responseBody map[string]interface{}
			err := json.Unmarshal(w.Body.Bytes(), &responseBody)

			assert.NoError(t, err)
			assert.Equal(t, tt.expectedBody, responseBody)
		})
	}
}
