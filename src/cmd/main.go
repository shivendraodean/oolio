package main

import (
	"log/slog"

	"oolio/api-ecommerce/src/internal/webapp"
)

func main() {
	container := webapp.NewContainer()
	router := webapp.SetupRouter(container)

	slog.Info("Starting server on :8080")
	if err := router.Run(":8080"); err != nil {
		slog.Error("Failed to start server", "error", err)
	}
}
