package models

import (
	"gorm.io/gorm"
)

type Inventory struct {
	gorm.Model
	VariantID uint    `gorm:"index"`
	Variant   Variant `gorm:"foreignKey:VariantID"`
	Quantity  int     `gorm:"default:0"`
}
