package webapp

import (
	"log/slog"

	"oolio/api-ecommerce/src/internal/repository"
	"oolio/api-ecommerce/src/internal/service"
	"oolio/api-ecommerce/src/internal/webapp/handler"
)

type Container struct {
	ProductHandler *handler.ProductHandler
	Logger         *slog.Logger
}

func NewContainer() *Container {
	// Webapp
	logger := InitLogger()

	// Repositories
	productRepo := repository.NewProductRepository()

	// Services
	productService := service.NewProductService(productRepo)

	// Initialize handlers with their dependencies
	productHandler := handler.NewProductHandler(productService)

	return &Container{
		ProductHandler: productHandler,
		Logger:         logger,
	}
}
