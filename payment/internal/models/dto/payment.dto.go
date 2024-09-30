package dto

import (
	"time"

	"github.com/VuThanhThien/golang-gorm-postgres/payment/internal/models"
)

type PaymentStatus string

const (
	PaymentStatusCompleted PaymentStatus = "COMPLETED"
	PaymentStatusPending   PaymentStatus = "PENDING"
	PaymentStatusFailed    PaymentStatus = "FAILED"
)

type Currency string

const (
	CurrencyUSD Currency = "USD"
	CurrencyEUR Currency = "EUR"
	CurrencyGBP Currency = "GBP"
	CurrencyJPY Currency = "JPY"
	CurrencyCNY Currency = "CNY"
	CurrencyVND Currency = "VND"
)

type PaymentMethod string

const (
	PaymentMethodCreditCard   PaymentMethod = "CREDIT_CARD"
	PaymentMethodPaypal       PaymentMethod = "PAYPAL"
	PaymentMethodBankTransfer PaymentMethod = "BANK_TRANSFER"
	PaymentMethodCash         PaymentMethod = "CASH"
)

type CreatePaymentDto struct {
	OrderID       uint          `json:"order_id" binding:"required" validate:"required" example:"1"`
	Amount        float64       `json:"amount" binding:"required" validate:"required" example:"100000"`
	Currency      Currency      `json:"currency" binding:"required" validate:"required" example:"USD"`
	Method        PaymentMethod `json:"method" binding:"required" validate:"required" example:"CREDIT_CARD"`
	Status        PaymentStatus `json:"status" binding:"required" validate:"required" example:"PENDING"`
	TransactionID string        `json:"transaction_id" binding:"required" validate:"required" example:"1234567890"`
}
type CreatePaymentInput struct {
	PaymentID     string        `json:"payment_id" binding:"required" validate:"required" example:"1234567890"`
	OrderID       uint          `json:"order_id" binding:"required" validate:"required" example:"1"`
	Amount        float64       `json:"amount" binding:"required" validate:"required" example:"100000"`
	Currency      Currency      `json:"currency" binding:"required" validate:"required" example:"USD"`
	Method        PaymentMethod `json:"method" binding:"required" validate:"required" example:"CREDIT_CARD"`
	Status        PaymentStatus `json:"status" binding:"required" validate:"required" example:"PENDING"`
	TransactionID string        `json:"transaction_id" binding:"required" validate:"required" example:"1234567890"`
	PaidAt        time.Time     `json:"paid_at" example:"2024-01-01T00:00:00Z"`
}

type FilterPaymentDto struct {
	PaymentID string        `json:"payment_id"`
	OrderID   uint          `json:"order_id"`
	Currency  Currency      `json:"currency"`
	Method    PaymentMethod `json:"method"`
	Status    PaymentStatus `json:"status"`
	PaidAt    time.Time     `json:"paid_at"`
}

type PaymentResponse struct {
	ID            uint          `json:"id,omitempty"`
	PaymentID     string        `json:"payment_id,omitempty"`
	OrderID       uint          `json:"order_id,omitempty"`
	Amount        float64       `json:"amount,omitempty"`
	Currency      Currency      `json:"currency,omitempty"`
	Method        PaymentMethod `json:"method,omitempty"`
	Status        PaymentStatus `json:"status,omitempty"`
	TransactionID string        `json:"transaction_id,omitempty"`
	PaidAt        time.Time     `json:"paid_at,omitempty"`
	CreatedAt     time.Time     `json:"created_at,omitempty"`
	UpdatedAt     time.Time     `json:"updated_at,omitempty"`
}

type PaginationResult struct {
	Data       []models.Payment
	TotalItems int64
	TotalPages int
	Page       int
	PageSize   int
}
