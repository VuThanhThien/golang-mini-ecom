package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name       string `gorm:"type:varchar(255);not null"`
	Email      string `gorm:"uniqueIndex;not null"`
	Password   string `gorm:"not null"`
	Role       string `gorm:"type:varchar(255);not null"`
	Provider   string `gorm:"not null"`
	Photo      string
	Verified   bool `gorm:"not null;default:true"`
	MerchantID uint `gorm:"default:null"`
}
