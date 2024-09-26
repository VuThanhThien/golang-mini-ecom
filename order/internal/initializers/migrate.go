package initializers

import (
	"github.com/VuThanhThien/golang-gorm-postgres/order/internal/models"
)

func Migrate() error {
	return DB.AutoMigrate(&models.Order{})
}
