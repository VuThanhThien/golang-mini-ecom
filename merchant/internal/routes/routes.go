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
	merchantController := controllers.NewMerchantController(merchantService)

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

	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productController := controllers.NewProductController(productService)

	productRoutes := router.Group("products")
	{
		productRoutes.POST("/", middleware.DeserializeUser(userGateway), productController.CreateProduct)
		productRoutes.GET("/:id", productController.GetProductByID)
		productRoutes.GET("/product-id/:productID", productController.GetProductByProductID)
		productRoutes.DELETE("/:id", productController.DeleteProduct)
	}

	categoryRepo := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryController := controllers.NewCategoryController(categoryService)

	categoryRoutes := router.Group("categories")
	{
		categoryRoutes.GET("/", categoryController.ListCategory)
		categoryRoutes.POST("/", middleware.DeserializeUser(userGateway), categoryController.CreateCategory)
		categoryRoutes.GET("/:id", categoryController.GetCategoryByID)
		categoryRoutes.DELETE("/:id", categoryController.DeleteCategory)
	}

	variantRepo := repositories.NewVariantRepository(db)
	variantService := services.NewVariantService(variantRepo)
	variantController := controllers.NewVariantController(variantService)

	variantRoutes := router.Group("variants")
	{
		variantRoutes.POST("/", middleware.DeserializeUser(userGateway), variantController.CreateVariant)
		variantRoutes.GET("/product-id/:productID", variantController.GetVariantByProductID)
		variantRoutes.GET("/variant-name/:variantName", variantController.GetVariantByVariantName)
		variantRoutes.DELETE("/:id", variantController.DeleteVariant)
	}

	inventoryRepo := repositories.NewInventoryRepository(db)
	inventoryService := services.NewInventoryService(inventoryRepo)
	inventoryController := controllers.NewInventoryController(inventoryService)

	inventoryRoutes := router.Group("inventory")
	{
		inventoryRoutes.POST("/", middleware.DeserializeUser(userGateway), inventoryController.CreateInventory)
		inventoryRoutes.GET("/:id", inventoryController.GetInventoryByID)
		inventoryRoutes.GET("/variant/:variant_id", inventoryController.GetInventoryByVariantID)
		inventoryRoutes.DELETE("/:id", inventoryController.DeleteInventory)
	}

	productImageRepo := repositories.NewProductImageRepository(db)
	productImageService := services.NewProductImageService(productImageRepo)
	productImageController := controllers.NewProductImageController(productImageService)

	productImageRoutes := router.Group("product-images")
	{
		productImageRoutes.POST("/", middleware.DeserializeUser(userGateway), productImageController.CreateProductImage)
		productImageRoutes.GET("/:id", productImageController.GetProductImageByID)
		productImageRoutes.PUT("/:id", productImageController.UpdateProductImage)
		productImageRoutes.DELETE("/:id", productImageController.DeleteProductImage)
	}

}
