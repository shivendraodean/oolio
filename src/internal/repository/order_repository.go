package repository

import (
	"database/sql"
	"encoding/json"
	"log"
	"oolio/api-ecommerce/src/internal/model"
)

type OrderRepository interface {
	CreateOrder(order model.Order) model.Order
}

type OrderRepositoryImpl struct {
	db *sql.DB
}

func NewOrderRepository() OrderRepository {
	db, err := GetDBConnection()
	if err != nil {
		log.Printf("Error connecting to database: %v", err)
		return &OrderRepositoryImpl{}
	}

	return &OrderRepositoryImpl{db: db}
}

func (r *OrderRepositoryImpl) CreateOrder(order model.Order) model.Order {
	if r.db == nil {
		log.Println("Database connection not established")
		return order
	}

	// Convert order items to JSON
	itemsJSON, err := json.Marshal(order.Items)
	if err != nil {
		log.Printf("Error marshaling order items: %v", err)
		return order
	}

	// Convert products to JSON
	productsJSON, err := json.Marshal(order.Products)
	if err != nil {
		log.Printf("Error marshaling products: %v", err)
		return order
	}

	// Insert order into database
	query := `INSERT INTO orders (items, products, coupon_code) 
			VALUES ($1, $2, $3) 
			RETURNING id`

	err = r.db.QueryRow(
		query,
		itemsJSON,
		productsJSON,
		order.CouponCode,
	).Scan(&order.ID)

	if err != nil {
		log.Printf("Error creating order: %v", err)
	}

	return order
}
