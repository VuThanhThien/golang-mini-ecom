package models

import (
	"gorm.io/gorm"
)

type Inventory struct {
	gorm.Model
	VariantID uint `gorm:"index"`
	Quantity  int  `gorm:"default:0"`
}
