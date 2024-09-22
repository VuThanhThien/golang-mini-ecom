package models

type Merchant struct {
	Base
	Name        string `gorm:"type:varchar(255);not null"`
	MerchantId  string `gorm:"uniqueIndex;not null"`
	Description string `gorm:"type:varchar(255)"`
}
