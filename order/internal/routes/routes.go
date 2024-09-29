package routes

import (
	"context"
	"net/http"

	"github.com/VuThanhThien/golang-gorm-postgres/order/internal/api/controllers"
	"github.com/VuThanhThien/golang-gorm-postgres/order/internal/api/repositories"
	"github.com/VuThanhThien/golang-gorm-postgres/order/internal/api/services"
	"github.com/VuThanhThien/golang-gorm-postgres/order/internal/gateway/grpc"
	"github.com/VuThanhThien/golang-gorm-postgres/order/internal/initializers"
	"github.com/VuThanhThien/golang-gorm-postgres/order/internal/middleware"
	"github.com/VuThanhThien/golang-gorm-postgres/order/internal/models/dto"
	"github.com/VuThanhThien/golang-gorm-postgres/order/internal/rabbit_handler"
	"github.com/VuThanhThien/golang-gorm-postgres/order/pkg/rabbitmq"
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

	//TODO: add refund publisher
	createOrderPublisher := rabbitmq.NewPublisher(
		context.Background(),
		&rabbitCfg,
		rabbitConn,
		log,
		rabbitmq.E_COM_EXCHANGE,
		"direct",
		rabbitmq.CREATE_ORDER_COMPLETED_QUEUE,
		rabbitmq.CREATE_ORDER_ROUTING_KEY,
	)
	userGateway := grpc.NewUserGateway(config.USER_GRPC_SERVER_HOST, config.USER_GRPC_SERVER_PORT)
	inventoryGateway := grpc.NewInventoryGateway(config.MERCHANT_GRPC_SERVER_HOST, config.MERCHANT_GRPC_SERVER_PORT)
	orderRepo := repositories.NewOrderRepository(db)
	orderItemRepo := repositories.NewItemRepository(db)
	orderService := services.NewOrderService(orderRepo, orderItemRepo, createOrderPublisher, inventoryGateway)
	orderController := controllers.NewOrderController(orderService)
	router.GET("/healthcheck", func(ctx *gin.Context) {
		message := "Welcome to Golang with Gorm and Postgres"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})
	orderRoutes := router.Group("orders")
	{
		orderRoutes.GET("/", middleware.DeserializeUser(userGateway), orderController.GetOrders)
		orderRoutes.GET("/:id", middleware.DeserializeUser(userGateway), orderController.GetOrder)
		orderRoutes.POST("/", middleware.DeserializeUser(userGateway), orderController.CreateOrder)
	}

	paymentDependencies := rabbit_handler.PaymentDependencies{
		Logger:       log,
		OrderService: orderService,
	}
	userConsumer := rabbitmq.NewConsumer(
		context.Background(),
		&rabbitCfg,
		rabbitConn,
		log,
		rabbit_handler.PaymentOrderCompleted,
		rabbitmq.E_COM_EXCHANGE,
		"direct",
		rabbitmq.PAYMENT_ORDER_COMPLETED_QUEUE,
		rabbitmq.PAYMENT_ORDER_COMPLETED_ROUTING_KEY,
	)
	go func() {
		err := userConsumer.ConsumeMessage(dto.PaymentResponse{}, &paymentDependencies)
		if err != nil {
			log.Error().Err(err).Msg("Consume message error")
		}
	}()
}
