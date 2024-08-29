package routes

import (
	"net/http"

	"github.com/VuThanhThien/golang-gorm-postgres/internal/api/controllers"
	"github.com/VuThanhThien/golang-gorm-postgres/internal/api/repositories"
	"github.com/VuThanhThien/golang-gorm-postgres/internal/api/services"
	"github.com/VuThanhThien/golang-gorm-postgres/internal/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(server *gin.Engine, db *gorm.DB) {
	router := server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		message := "Welcome to Golang with Gorm and Postgres"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	authRepo := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepo)
	authController := controllers.NewAuthController(db, authService)

	authRouters := router.Group("/auth")
	authRouters.POST("/register", authController.SignUpUser)
	authRouters.POST("/login", authController.SignInUser)
	authRouters.GET("/refresh", authController.RefreshAccessToken)
	authRouters.GET("/logout", middleware.DeserializeUser(), authController.LogoutUser)

	postController := controllers.NewPostController(db)
	postRoutes := router.Group("posts")
	postRoutes.Use(middleware.DeserializeUser())
	postRoutes.POST("/", postController.CreatePost)
	postRoutes.GET("/", postController.FindPosts)
	postRoutes.PUT("/:postId", postController.UpdatePost)
	postRoutes.GET("/:postId", postController.FindPostById)
	postRoutes.DELETE("/:postId", postController.DeletePost)

	userController := controllers.NewUserController(db)
	userRoutes := router.Group("users")
	userRoutes.GET("/me", middleware.DeserializeUser(), userController.GetMe)

}
