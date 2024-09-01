package models

import (
	"gorm.io/gorm"
	"time"

	"github.com/google/uuid"
)

type User struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name      string    `gorm:"type:varchar(255);not null"`
	Email     string    `gorm:"uniqueIndex;not null"`
	Password  string    `gorm:"not null"`
	Role      string    `gorm:"type:varchar(255);not null"`
	Provider  string    `gorm:"not null"`
	Photo     string
	Verified  bool      `gorm:"not null;default:true"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
	Posts     []Post    `gorm:"foreignKey:UserID" json:"posts"`
}
