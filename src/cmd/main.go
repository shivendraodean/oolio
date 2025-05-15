package main

import (
	"log"

	"oolio/api-ecommerce/src/internal/webapp"
)

func main() {
	container := webapp.NewContainer()
	router := webapp.SetupRouter(container)

	log.Println("Starting server on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
