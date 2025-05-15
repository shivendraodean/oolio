package webapp

import (
	"oolio/api-ecommerce/src/internal/handler"
	"oolio/api-ecommerce/src/internal/repository"
	"oolio/api-ecommerce/src/internal/service"
)

type Container struct {
	ProductHandler *handler.ProductHandler
}

func NewContainer() *Container {
	// Repositories
	productRepo := repository.NewProductRepository()

	// Services
	productService := service.NewProductService(productRepo)

	// Initialize handlers with their dependencies
	productHandler := handler.NewProductHandler(productService)

	return &Container{
		ProductHandler: productHandler,
	}
}
