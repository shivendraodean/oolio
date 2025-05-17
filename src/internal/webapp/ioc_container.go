package webapp

import (
	"log/slog"

	"oolio/api-ecommerce/src/internal/repository"
	"oolio/api-ecommerce/src/internal/service"
	"oolio/api-ecommerce/src/internal/webapp/handler"
)

type Container struct {
	ProductHandler *handler.ProductHandler
	OrderHandler   *handler.OrderHandler
	Logger         *slog.Logger
}

func NewContainer() *Container {
	// Webapp
	logger := InitLogger()

	// Repositories
	productRepo := repository.NewProductRepository()
	orderRepo := repository.NewOrderRepository()
	couponRepo := repository.NewCouponRepository()

	// Services
	productService := service.NewProductService(productRepo)
	orderService := service.NewOrderService(orderRepo, productRepo, couponRepo)

	// Initialize handlers with their dependencies
	productHandler := handler.NewProductHandler(productService)
	orderHandler := handler.NewOrderHandler(orderService)

	return &Container{
		ProductHandler: productHandler,
		OrderHandler:   orderHandler,
		Logger:         logger,
	}
}
