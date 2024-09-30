package dto

import "time"

type Item struct {
	ID         uint      `json:"id" `
	OrderID    uint      `json:"order_id"`
	ProductID  uint      `json:"product_id"`
	VariantID  uint      `json:"variant_id"`
	Name       string    `json:"name"`
	Quantity   int       `json:"quantity"`
	Price      float64   `json:"price"`
	TotalPrice float64   `json:"total_price"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  time.Time `json:"deleted_at"`
}

type CreatedOrder struct {
	ID          uint    `json:"id"`
	OrderID     string  `json:"order_id"`
	UserID      uint    `json:"user_id"`
	PaymentID   uint    `json:"payment_id"`
	Status      string  `json:"status"`
	TotalAmount float64 `json:"total_amount"`
	Items       []Item  `json:"items"`
	PlacedAt    string  `json:"placed_at"`
}
