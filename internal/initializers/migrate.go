package initializers

import (
	"log"

	"github.com/VuThanhThien/golang-gorm-postgres/internal/models"
)

func init() {
	config, err := LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	ConnectDB(&config)
}

func Migrate() error {
	DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	return DB.AutoMigrate(&models.User{}, &models.Post{})
}
