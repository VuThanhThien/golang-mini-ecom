package dto

import (
	"time"

	"github.com/VuThanhThien/golang-gorm-postgres/order/internal/models"
)

type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "PENDING"
	OrderStatusCompleted OrderStatus = "COMPLETED"
	OrderStatusCancelled OrderStatus = "CANCELLED"
	OrderStatusFailed    OrderStatus = "FAILED"
)

func (s OrderStatus) IsValid() bool {
	switch s {
	case OrderStatusPending, OrderStatusCompleted, OrderStatusCancelled, OrderStatusFailed:
		return true
	}
	return false
}

func (s OrderStatus) String() string {
	return string(s)
}

type ItemDto struct {
	ProductID  uint    `json:"product_id" example:"1"`
	VariantID  uint    `json:"variant_id" example:"1"`
	Name       string  `json:"name" example:"Quần"`
	Quantity   int     `json:"quantity" example:"1"`
	Price      float64 `json:"price" example:"100000"`
	TotalPrice float64 `json:"total_price" example:"100000"`
}

type GetOrderRequestDto struct {
	OrderID    uint        `json:"order_id"`
	UserID     uint        `json:"user_id"`
	MerchantID uint        `json:"merchant_id"`
	PaymentID  uint        `json:"payment_id"`
	Status     OrderStatus `json:"status"`
}

type GetOrderResponseDto struct {
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
	DeletedAt   time.Time     `json:"deleted_at"`
	OrderID     uint          `json:"order_id"`
	UserID      uint          `json:"user_id"`
	MerchantID  uint          `json:"merchant_id"`
	PaymentID   uint          `json:"payment_id"`
	Status      OrderStatus   `json:"status"`
	TotalAmount float64       `json:"total_amount"`
	Items       []models.Item `json:"items"`
}

type CreateOrderRequestDto struct {
	OrderID     string    `json:"order_id" validate:"required" example:"ORDER-123456"`
	UserID      uint      `json:"user_id" validate:"required" example:"1"`
	TotalAmount float64   `json:"total_amount" validate:"required" example:"100000"`
	Items       []ItemDto `json:"items" validate:"required"`
}

type CreateOrderResponseDto struct {
	OrderID     uint          `json:"order_id"`
	UserID      uint          `json:"user_id"`
	TotalAmount float64       `json:"total_amount"`
	Items       []models.Item `json:"items"`
	Status      OrderStatus   `json:"status"`
	PlacedAt    time.Time     `json:"placed_at"`
}

type UpdateOrderRequestDto struct {
	ID     uint        `json:"id" validate:"required" example:"1"`
	Status OrderStatus `json:"status" validate:"required" example:"PAID"`
}

type PaginationResult struct {
	Data       []models.Order `json:"data"`
	TotalItems int64          `json:"total_items"`
	TotalPages int            `json:"total_pages"`
	Page       int            `json:"page"`
	PageSize   int            `json:"page_size"`
}
