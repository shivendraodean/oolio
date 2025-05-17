package model

type OrderItem struct {
	ProductID int64 `json:"productId"`
	Quantity  int   `json:"quantity"`
}

type Order struct {
	ID         string      `json:"id"`
	Items      []OrderItem `json:"items"`
	Products   []Product   `json:"products"`
	CouponCode string      `json:"couponCode,omitempty"`
}
