package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	OrderID     string    `gorm:"uniqueIndex;not null" json:"order_id"`
	UserID      uint      `gorm:"not null" json:"user_id"`
	PaymentID   uint      `gorm:"default:null" json:"payment_id"`
	Status      string    `gorm:"type:varchar(50);not null" json:"status"`
	TotalAmount float64   `gorm:"type:decimal(10,2);not null" json:"total_amount"`
	Items       []Item    `gorm:"foreignKey:OrderID;references:ID" json:"items"`
	PlacedAt    time.Time `gorm:"not null; default:now()"`
}

type Item struct {
	gorm.Model
	OrderID    uint    `gorm:"not null" json:"order_id"`
	ProductID  uint    `gorm:"not null" json:"product_id"`
	VariantID  uint    `gorm:"" json:"variant_id"`
	Name       string  `gorm:"type:varchar(255);not null" json:"name"`
	Quantity   int     `gorm:"not null" json:"quantity"`
	Price      float64 `gorm:"type:decimal(10,2);not null" json:"price"`
	TotalPrice float64 `gorm:"type:decimal(10,2);not null" json:"total_price"`
}
