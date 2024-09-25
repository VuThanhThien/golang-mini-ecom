package models

import (
	"gorm.io/gorm"
)

type Variant struct {
	gorm.Model
	ProductID   uint    `gorm:"index"`
	VariantName string  `gorm:"type:varchar(255)"`
	Description string  `gorm:"type:text"`
	Price       float64 `gorm:"type:decimal(10,2)"`
}
