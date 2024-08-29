package initializers

import (
	"fmt"
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

func Migrate() {
	DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	DB.AutoMigrate(&models.User{}, &models.Post{})
	fmt.Println("👍 Migration complete")
}
