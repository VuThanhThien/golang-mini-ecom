package routes

import (
	"context"
	"net/http"

	"github.com/VuThanhThien/golang-gorm-postgres/payment/internal/api/controllers"
	"github.com/VuThanhThien/golang-gorm-postgres/payment/internal/api/repositories"
	"github.com/VuThanhThien/golang-gorm-postgres/payment/internal/api/services"
	"github.com/VuThanhThien/golang-gorm-postgres/payment/internal/initializers"
	"github.com/VuThanhThien/golang-gorm-postgres/payment/pkg/rabbitmq"
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
	paymentOrderCompletedPublisher := rabbitmq.NewPublisher(
		context.Background(),
		&rabbitCfg,
		rabbitConn,
		log,
		rabbitmq.E_COM_EXCHANGE,
		"direct",
		rabbitmq.PAYMENT_ORDER_COMPLETED_QUEUE,
		rabbitmq.PAYMENT_ORDER_COMPLETED_ROUTING_KEY,
	)
	paymentRepo := repositories.NewPaymentRepository(db)
	paymentService := services.NewPaymentService(paymentRepo, paymentOrderCompletedPublisher)
	paymentController := controllers.NewPaymentController(paymentService)
	router.GET("/healthcheck", func(ctx *gin.Context) {
		message := "Welcome to Golang with Gorm and Postgres"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})
	paymentRoutes := router.Group("payments")
	{
		paymentRoutes.GET("/", paymentController.ListPayments)
		paymentRoutes.GET("/:id", paymentController.ReadPayment)
		paymentRoutes.POST("/", paymentController.CreatePayment)
	}

}
