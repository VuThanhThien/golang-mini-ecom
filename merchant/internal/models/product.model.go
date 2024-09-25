package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string   `gorm:"type:varchar(255);not null"`
	Description string   `gorm:"type:text"`
	Price       float64  `gorm:"type:decimal(10,2);not null"`
	SKU         string   `gorm:"uniqueIndex;not null"`
	MerchantID  uint     `gorm:"default:null"`
	Merchant    Merchant `gorm:"foreignKey:MerchantID"`
	CategoryID  uint     `gorm:"default:null"`
	Category    Category `gorm:"foreignKey:CategoryID"`
}
