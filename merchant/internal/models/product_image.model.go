package models

import (
	"gorm.io/gorm"
)

type ProductImage struct {
	gorm.Model
	ProductID uint   `gorm:"index"`
	ImageURL  string `gorm:"type:text"`
	IsPrimary bool   `gorm:"default:false"`
}
