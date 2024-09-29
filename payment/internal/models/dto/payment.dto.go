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
)

type PaymentMethod string

const (
	PaymentMethodCreditCard   PaymentMethod = "CREDIT_CARD"
	PaymentMethodPaypal       PaymentMethod = "PAYPAL"
	PaymentMethodBankTransfer PaymentMethod = "BANK_TRANSFER"
	PaymentMethodCash         PaymentMethod = "CASH"
)

type CreatePaymentDto struct {
	OrderID       uint          `json:"order_id" binding:"required"`
	Amount        float64       `json:"amount" binding:"required"`
	Currency      Currency      `json:"currency" binding:"required"`
	Method        PaymentMethod `json:"method" binding:"required"`
	Status        PaymentStatus `json:"status" binding:"required"`
	TransactionID string        `json:"transaction_id"`
}
type CreatePaymentInput struct {
	PaymentID     string        `json:"payment_id" binding:"required"`
	OrderID       uint          `json:"order_id" binding:"required"`
	Amount        float64       `json:"amount" binding:"required"`
	Currency      Currency      `json:"currency" binding:"required"`
	Method        PaymentMethod `json:"method" binding:"required"`
	Status        PaymentStatus `json:"status" binding:"required"`
	TransactionID string        `json:"transaction_id"`
	PaidAt        time.Time     `json:"paid_at"`
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
