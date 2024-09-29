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
	"github.com/VuThanhThien/golang-gorm-postgres/merchant/internal/models/dto"
	"github.com/VuThanhThien/golang-gorm-postgres/merchant/internal/rabbit_handler"
	"github.com/VuThanhThien/golang-gorm-postgres/merchant/pkg/rabbitmq"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/streadway/amqp"
	"gorm.io/gorm"
)

func SetupRoutes(server *gin.Engine, db *gorm.DB, rabbitConn *amqp.Connection, log zerolog.Logger, config *initializers.Config, ctx context.Context) {
	router := server.Group("/api")
	rabbitCfg := rabbitmq.RabbitMQConfig{
		Host:     config.AMQP_SERVER_HOST,
		Port:     config.AMQP_SERVER_PORT,
		User:     config.AMQP_SERVER_USER,
		Password: config.AMQP_SERVER_PASSWORD,
	}

	createOrderPublisher := rabbitmq.NewPublisher(
		ctx,
		&rabbitCfg,
		rabbitConn,
		log,
		rabbitmq.E_COM_EXCHANGE,
		"direct",
		rabbitmq.CREATE_ORDER_COMPLETED_QUEUE,
	)
	userGateway := grpc.New(config.USER_GRPC_SERVER_HOST, config.USER_GRPC_SERVER_PORT)

	merchantRepo := repositories.NewMerchantRepository(db)
	merchantService := services.NewMerchantService(merchantRepo, createOrderPublisher)
	merchantController := controllers.NewMerchantController(merchantService)
	router.GET("/healthcheck", func(ctx *gin.Context) {
		message := "Welcome to Golang with Gorm and Postgres"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})
	merchantRoutes := router.Group("merchants")
	{
		merchantRoutes.GET("/:id", middleware.DeserializeUser(userGateway), merchantController.GetMerchant)
		merchantRoutes.GET("/merchant-id/:merchantID", middleware.DeserializeUser(userGateway), merchantController.GetMerchantByMerchantID)
		merchantRoutes.POST("/", middleware.DeserializeUser(userGateway), merchantController.CreateMerchant)
	}

	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productController := controllers.NewProductController(productService, merchantService)

	productRoutes := router.Group("products")
	{
		productRoutes.GET("/", middleware.DeserializeUser(userGateway), productController.FilterProductsWithPagination)
		productRoutes.POST("/", middleware.DeserializeUser(userGateway), productController.CreateProduct)
		productRoutes.GET("/:id", middleware.DeserializeUser(userGateway), productController.GetProductByID)
		productRoutes.GET("/product-id/:productID", middleware.DeserializeUser(userGateway), productController.GetProductByProductID)
		productRoutes.DELETE("/:id", middleware.DeserializeUser(userGateway), productController.DeleteProduct)
	}

	categoryRepo := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryController := controllers.NewCategoryController(categoryService)

	categoryRoutes := router.Group("categories")
	{
		categoryRoutes.GET("/", middleware.DeserializeUser(userGateway), categoryController.ListCategory)
		categoryRoutes.POST("/", middleware.DeserializeUser(userGateway), categoryController.CreateCategory)
		categoryRoutes.GET("/:id", middleware.DeserializeUser(userGateway), categoryController.GetCategoryByID)
		categoryRoutes.DELETE("/:id", middleware.DeserializeUser(userGateway), categoryController.DeleteCategory)
	}

	variantRepo := repositories.NewVariantRepository(db)
	variantService := services.NewVariantService(variantRepo)
	variantController := controllers.NewVariantController(variantService)

	variantRoutes := router.Group("variants")
	{
		variantRoutes.POST("/", middleware.DeserializeUser(userGateway), variantController.CreateVariant)
		variantRoutes.GET("/:id", middleware.DeserializeUser(userGateway), variantController.GetVariantById)
		variantRoutes.GET("/product-id/:id", middleware.DeserializeUser(userGateway), variantController.GetVariantByProductID)
		variantRoutes.GET("/variant-name/:variantName", middleware.DeserializeUser(userGateway), variantController.GetVariantByVariantName)
		variantRoutes.DELETE("/:id", middleware.DeserializeUser(userGateway), variantController.DeleteVariant)
	}

	inventoryRepo := repositories.NewInventoryRepository(db)
	inventoryService := services.NewInventoryService(inventoryRepo)
	inventoryController := controllers.NewInventoryController(inventoryService)

	inventoryRoutes := router.Group("inventory")
	{
		inventoryRoutes.POST("/", middleware.DeserializeUser(userGateway), inventoryController.CreateInventory)
		inventoryRoutes.GET("/:id", inventoryController.GetInventoryByID)
		inventoryRoutes.GET("/variant/:id", middleware.DeserializeUser(userGateway), inventoryController.GetInventoryByVariantID)
		inventoryRoutes.DELETE("/:id", middleware.DeserializeUser(userGateway), inventoryController.DeleteInventory)
		inventoryRoutes.POST("/order-succeed", middleware.DeserializeUser(userGateway), inventoryController.DeductQuantity)
		inventoryRoutes.POST("/order-failed", middleware.DeserializeUser(userGateway), inventoryController.RefundQuantity)
	}

	productImageRepo := repositories.NewProductImageRepository(db)
	productImageService := services.NewProductImageService(productImageRepo)
	productImageController := controllers.NewProductImageController(productImageService)

	productImageRoutes := router.Group("product-images")
	{
		productImageRoutes.POST("/", middleware.DeserializeUser(userGateway), productImageController.CreateProductImage)
		productImageRoutes.GET("/:id", middleware.DeserializeUser(userGateway), productImageController.GetProductImageByID)
		productImageRoutes.PUT("/:id", middleware.DeserializeUser(userGateway), productImageController.UpdateProductImage)
		productImageRoutes.DELETE("/:id", middleware.DeserializeUser(userGateway), productImageController.DeleteProductImage)
	}

	OrderDependencies := rabbit_handler.OrderDependencies{
		Logger:           log,
		InventoryService: inventoryService,
	}
	userConsumer := rabbitmq.NewConsumer(
		ctx,
		&rabbitCfg,
		rabbitConn,
		log,
		rabbit_handler.CreateOrderSucceed,
		rabbitmq.E_COM_EXCHANGE,
		"direct",
		rabbitmq.CREATE_ORDER_COMPLETED_QUEUE,
		rabbitmq.CREATE_ORDER_ROUTING_KEY,
	)
	go func() {
		err := userConsumer.ConsumeMessage(dto.CreateOrder{}, &OrderDependencies)
		if err != nil {
			log.Error().Err(err).Msg("Consume message error")
		}
	}()

}
