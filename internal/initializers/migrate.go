package initializers

import (
	"github.com/VuThanhThien/golang-gorm-postgres/internal/models"
)

func Migrate() error {
	DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	return DB.AutoMigrate(&models.User{}, &models.Post{})
}
