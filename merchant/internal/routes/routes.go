package routes

import (
	"net/http"

	"github.com/VuThanhThien/golang-gorm-postgres/merchant/internal/api/controllers"
	"github.com/VuThanhThien/golang-gorm-postgres/merchant/internal/api/repositories"
	"github.com/VuThanhThien/golang-gorm-postgres/merchant/internal/api/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(server *gin.Engine, db *gorm.DB) {
	router := server.Group("/api")

	merchantRepo := repositories.NewMerchantRepository(db)
	merchantService := services.NewMerchantService(merchantRepo)
	merchantController := controllers.NewMerchantController(db, merchantService)

	merchantRoutes := router.Group("merchants")
	{
		router.GET("/healthcheck", func(ctx *gin.Context) {
			message := "Welcome to Golang with Gorm and Postgres"
			ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
		})
		merchantRoutes.GET("/:id", merchantController.GetMerchant)
		merchantRoutes.GET("/merchant-id/:merchantID", merchantController.GetMerchantByMerchantID)
		merchantRoutes.POST("/", merchantController.CreateMerchant)
	}

}
