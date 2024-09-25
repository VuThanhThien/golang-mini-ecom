package models

import (
	"gorm.io/gorm"
)

type ProductImage struct {
	gorm.Model
	ProductID uint   `gorm:"index"`
	ImageURL  string `gorm:"type:varchar(255)"`
	IsPrimary bool   `gorm:"default:false"`
}
