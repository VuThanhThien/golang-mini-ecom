package routes

import (
	"context"
	"net/http"

	"github.com/VuThanhThien/golang-gorm-postgres/merchant/internal/api/controllers"
	"github.com/VuThanhThien/golang-gorm-postgres/merchant/internal/api/repositories"
	"github.com/VuThanhThien/golang-gorm-postgres/merchant/internal/api/services"
	"github.com/VuThanhThien/golang-gorm-postgres/merchant/internal/gateway/user/grpc"
	"github.com/VuThanhThien/golang-gorm-postgres/merchant/internal/initializers"
	"github.com/VuThanhThien/golang-gorm-postgres/merchant/internal/middleware"
	"github.com/VuThanhThien/golang-gorm-postgres/merchant/pkg/rabbitmq"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/streadway/amqp"
	"gorm.io/gorm"
)

func SetupRoutes(server *gin.Engine, db *gorm.DB, rabbitConn *amqp.Connection, log zerolog.Logger, config *initializers.Config) {
	router := server.Group("/api")
	rabbitCfg := rabbitmq.RabbitMQConfig{
		Host:     config.AMQP_SERVER_HOST,
		Port:     config.AMQP_SERVER_PORT,
		User:     config.AMQP_SERVER_USER,
		Password: config.AMQP_SERVER_PASSWORD,
	}

	createOrderPublisher := rabbitmq.NewPublisher(
		context.Background(),
		&rabbitCfg,
		rabbitConn,
		log,
		rabbitmq.E_COM_EXCHANGE,
		"direct",
		rabbitmq.PAYMENT_ORDER_COMPLETED_QUEUE,
	)
	userGateway := grpc.New(config.USER_GRPC_SERVER_HOST, config.USER_GRPC_SERVER_PORT)

	merchantRepo := repositories.NewMerchantRepository(db)
	merchantService := services.NewMerchantService(merchantRepo, createOrderPublisher)
	merchantController := controllers.NewMerchantController(db, merchantService)

	merchantRoutes := router.Group("merchants")
	{
		router.GET("/healthcheck", func(ctx *gin.Context) {
			message := "Welcome to Golang with Gorm and Postgres"
			ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
		})
		merchantRoutes.GET("/:id", merchantController.GetMerchant)
		merchantRoutes.GET("/merchant-id/:merchantID", merchantController.GetMerchantByMerchantID)
		merchantRoutes.POST("/", middleware.DeserializeUser(userGateway), merchantController.CreateMerchant)
	}

}
