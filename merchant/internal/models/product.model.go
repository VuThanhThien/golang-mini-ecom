package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string    `gorm:"type:varchar(255);not null"`
	Description string    `gorm:"type:text"`
	Price       float64   `gorm:"type:decimal(10,2);not null"`
	Stock       int       `gorm:"not null"`
	SKU         string    `gorm:"uniqueIndex;not null"`
	MerchantID  uint      `gorm:"default:null"`
	Merchant    Merchant  `gorm:"foreignKey:MerchantID"`
	CategoryID  uint      `gorm:"not null"`
	Category    Category  `gorm:"foreignKey:CategoryID"`
	Images      []Image   `gorm:"foreignKey:ProductID"`
	Variants    []Variant `gorm:"foreignKey:ProductID"`
}

type Category struct {
	gorm.Model
	Name        string     `gorm:"type:varchar(255);not null"`
	Description string     `gorm:"type:text"`
	ParentID    *uint      `gorm:"index"`
	Parent      *Category  `gorm:"foreignKey:ParentID"`
	Children    []Category `gorm:"foreignKey:ParentID"`
}

type Image struct {
	gorm.Model
	URL       string `gorm:"type:varchar(255);not null"`
	ProductID uint   `gorm:"not null"`
}

type Variant struct {
	gorm.Model
	Name      string  `gorm:"type:varchar(255);not null"`
	SKU       string  `gorm:"uniqueIndex;not null"`
	Price     float64 `gorm:"type:decimal(10,2);not null"`
	Stock     int     `gorm:"not null"`
	ProductID uint    `gorm:"not null"`
}
