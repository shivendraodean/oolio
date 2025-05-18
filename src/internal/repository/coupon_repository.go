package repository

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type CouponRepository interface {
	SearchCoupon(code string) int
}

type CouponRepositoryImpl struct {
	db *sql.DB
}

func NewCouponRepository() CouponRepository {
	db, err := GetDBConnection()
	if err != nil {
		log.Printf("Error connecting to database: %v", err)
		return &CouponRepositoryImpl{}
	}

	return &CouponRepositoryImpl{db: db}
}

func (r *CouponRepositoryImpl) SearchCoupon(code string) int {
	var count int
	query := "SELECT COUNT(*) FROM couponcodes WHERE code = $1"
	err := r.db.QueryRow(query, code).Scan(&count)
	if err != nil {
		log.Printf("Error searching for coupon code: %v", err)
		return 0
	}

	return count
}
