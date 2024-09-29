package models

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	PaymentID     string    `gorm:"uniqueIndex;not null" json:"payment_id"`
	OrderID       uint      `gorm:"not null" json:"order_id"`
	Amount        float64   `gorm:"type:decimal(10,2);not null" json:"amount"`
	Currency      string    `gorm:"type:varchar(3);not null" json:"currency"`
	Method        string    `gorm:"type:varchar(50);not null" json:"method"`
	Status        string    `gorm:"type:varchar(50);not null" json:"status"`
	TransactionID string    `gorm:"type:varchar(255)" json:"transaction_id"`
	PaidAt        time.Time `gorm:"" json:"paid_at"`
}
