package dto

import "time"

type PaymentStatus string

const (
	PaymentStatusCompleted PaymentStatus = "COMPLETED"
	PaymentStatusPending   PaymentStatus = "PENDING"
	PaymentStatusFailed    PaymentStatus = "FAILED"
)

type PaymentResponse struct {
	ID            uint          `json:"id,omitempty"`
	PaymentID     string        `json:"payment_id,omitempty"`
	OrderID       uint          `json:"order_id,omitempty"`
	Amount        float64       `json:"amount,omitempty"`
	Currency      string        `json:"currency,omitempty"`
	Method        string        `json:"method,omitempty"`
	Status        PaymentStatus `json:"status,omitempty"`
	TransactionID string        `json:"transaction_id,omitempty"`
	PaidAt        time.Time     `json:"paid_at,omitempty"`
	CreatedAt     time.Time     `json:"created_at,omitempty"`
	UpdatedAt     time.Time     `json:"updated_at,omitempty"`
}
