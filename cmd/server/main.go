package main

import (
	"errors"
	"fmt"
	"golang.org/x/net/context"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/VuThanhThien/golang-gorm-postgres/docs"
	"github.com/VuThanhThien/golang-gorm-postgres/internal/initializers"
	"github.com/VuThanhThien/golang-gorm-postgres/internal/middleware"
	"github.com/VuThanhThien/golang-gorm-postgres/internal/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const OperationIDKey = "X-Request-Id"

var (
	server *gin.Engine
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
	initializers.InitRedis(&config)
	if config.EnableAutoMigrate == "true" {
		if err := initializers.Migrate(); err != nil {
			log.Fatal("Failed to run database migrations", err)
		}
	}

	server = gin.Default()
}

//	@title						Swagger Example API
//	@version					1.0
//	@description				This is a sample server golang server.
//	@termsOfService				http://swagger.io/terms/
//	@contact.name				API Support
//	@contact.url				http://www.swagger.io/support
//	@contact.email				support@swagger.io
//	@license.name				Apache 2.0
//	@license.url				http://www.apache.org/licenses/LICENSE-2.0.html
//	@host						localhost:8000
//	@BasePath					/api
//	@securityDefinitions.basic	BasicAuth
//	@securityDefinitions.apikey	Bearer
//	@in							header
//	@name						Authorization
//	@description				Type "Bearer" followed by a space and JWT token.

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:" + config.ServerPort, config.ClientOrigin}
	corsConfig.AllowCredentials = true

	// To implement a graceful shutdown while using Gin, you should create a http.Server
	// instance yourself and avoid using server.Run()
	srv := &http.Server{
		Addr:    ":" + config.ServerPort,
		Handler: server.Handler(),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Listen: %s\n", err)
		}
	}()

	server.Use(cors.New(corsConfig))
	server.Use(middleware.RequestIDMiddleware())
	server.Use(middleware.LoggingMiddleware())
	server.Use(gzip.Gzip(gzip.DefaultCompression))

	// add timestamp to name to avoid overwrite this log
	f, _ := os.Create("tmp/gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	server.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("[%s] %s - [%s] \"%s: %s %s %d\" %s %s\n",
			param.TimeStamp.Format(time.DateTime),
			param.ClientIP,
			param.Keys[OperationIDKey],
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.ErrorMessage,
		)
	}))

	server.LoadHTMLGlob("./public/html/*")
	server.Static("/public", "./public")
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	server.NoRoute(func(c *gin.Context) {
		c.HTML(404, "404.html", gin.H{})
	})

	routes.SetupRoutes(server, initializers.DB)
	fmt.Printf("🚀 ~running on: http://localhost:%s/swagger/index.html 🚀 \n", config.ServerPort)

	// Wait for interrupt signal to gracefully shut down the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

}
