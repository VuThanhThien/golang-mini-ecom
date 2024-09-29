package initializers

import (
	"github.com/VuThanhThien/golang-gorm-postgres/payment/internal/models"
)

func Migrate() error {
	return DB.AutoMigrate(&models.Payment{})
}
