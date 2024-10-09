package routes

import (
	"net/http"

	"github.com/VuThanhThien/golang-gorm-postgres/authentication/internal/api/controllers"
	"github.com/VuThanhThien/golang-gorm-postgres/authentication/internal/api/repositories"
	"github.com/VuThanhThien/golang-gorm-postgres/authentication/internal/api/services"
	"github.com/VuThanhThien/golang-gorm-postgres/authentication/internal/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(server *gin.Engine, db *gorm.DB) {
	router := server.Group("/api")

	authRepo := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepo)
	authController := controllers.NewAuthController(db, authService)

	authRouters := router.Group("/auth")
	{
		authRouters.GET("/healthcheck", func(ctx *gin.Context) {
			message := "Service authentication is running..."
			ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
		})
		authRouters.POST("/register", authController.SignUpUser)
		authRouters.POST("/login", authController.SignInUser)
		authRouters.GET("/refresh", authController.RefreshAccessToken)
		authRouters.GET("/logout", middleware.DeserializeUser(), authController.LogoutUser)
	}

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(db, userService)

	userRoutes := router.Group("users")
	{
		userRoutes.GET("/list", middleware.DeserializeUser(), middleware.RequireRole(middleware.ADMIN), userController.ListUsers)
		userRoutes.GET("/me", middleware.DeserializeUser(), userController.GetMe)
		userRoutes.GET("/:id", middleware.DeserializeUser(), middleware.RequireRole(middleware.ADMIN), userController.GetUser)
	}

}
