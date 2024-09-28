package models

import (
	"gorm.io/gorm"
)

type Merchant struct {
	gorm.Model
	Name         string `gorm:"type:varchar(255);not null"`
	MerchantCode string `gorm:"uniqueIndex;not null"`
	Description  string `gorm:"type:varchar(255)"`
	UserID       uint   `gorm:"not null"`
}
