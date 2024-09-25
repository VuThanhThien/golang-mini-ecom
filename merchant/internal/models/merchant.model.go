package models

import (
	"gorm.io/gorm"
)

type Merchant struct {
	gorm.Model
	Name        string `gorm:"type:varchar(255);not null"`
	MerchantId  string `gorm:"uniqueIndex;not null"`
	Description string `gorm:"type:varchar(255)"`
	UserId      uint   `gorm:"not null"`
}
