package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name        string    `gorm:"type:varchar(255);not null"`
	Description string    `gorm:"type:text"`
	ParentID    uint      `gorm:"index;default:null"`
	Parent      *Category `gorm:"foreignKey:ParentID"`
}
