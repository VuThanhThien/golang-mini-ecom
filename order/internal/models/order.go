package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	OrderID     string    `gorm:"uniqueIndex;not null"`
	UserID      uint      `gorm:"not null"`
	PaymentID   uint      `gorm:"default:null"`
	Status      string    `gorm:"type:varchar(50);not null"`
	TotalAmount float64   `gorm:"type:decimal(10,2);not null"`
	Items       []Item    `gorm:"foreignKey:OrderID"`
	PlacedAt    time.Time `gorm:"not null; default:now()"`
}

type Item struct {
	gorm.Model
	OrderID    string  `gorm:"not null"`
	ProductID  uint    `gorm:"not null"`
	VariantID  uint    `gorm:""`
	Name       string  `gorm:"type:varchar(255);not null"`
	Quantity   int     `gorm:"not null"`
	Price      float64 `gorm:"type:decimal(10,2);not null"`
	TotalPrice float64 `gorm:"type:decimal(10,2);not null"`
}
