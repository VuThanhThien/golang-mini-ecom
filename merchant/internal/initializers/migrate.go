package initializers

import (
	"github.com/VuThanhThien/golang-gorm-postgres/merchant/internal/models"
)

func Migrate() error {
	return DB.AutoMigrate(&models.Merchant{})
}
