package controllers

import (
	"net/http"
	"strings"

	"github.com/VuThanhThien/golang-gorm-postgres/authentication/internal/api/services"
	"github.com/VuThanhThien/golang-gorm-postgres/authentication/internal/initializers"
	"github.com/VuThanhThien/golang-gorm-postgres/authentication/internal/models"
	"github.com/VuThanhThien/golang-gorm-postgres/authentication/internal/models/dto"
	"github.com/VuThanhThien/golang-gorm-postgres/authentication/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthController struct {
	DB      *gorm.DB
	service services.AuthServiceInterface
}

func NewAuthController(DB *gorm.DB, service services.AuthServiceInterface) AuthController {
	return AuthController{DB, service}
}

// SignUpUser godoc
//
//	@Summary		SignUpUser
//	@Description	SignUpUser
//	@Tags			auth
//	@Param			payload	body		dto.SignUpInput	true	"Register payload"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	dto.UserResponse
//	@Router			/auth/register [post]
func (ac *AuthController) SignUpUser(ctx *gin.Context) {
	var payload *dto.SignUpInput

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	if payload.Password != payload.PasswordConfirm {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Passwords do not match"})
		return
	}

	hashedPassword, err := utils.HashPassword(payload.Password)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
		return
	}

	newUser := models.User{
		Name:     payload.Name,
		Email:    strings.ToLower(payload.Email),
		Password: hashedPassword,
		Role:     "user",
		Verified: true,
		Photo:    payload.Photo,
		Provider: "local",
	}

	result := ac.service.SignUpUser(&newUser)

	if result != nil && strings.Contains(result.Error(), "duplicate key value violates unique") {
		ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "User with that email already exists"})
		return
	} else if result != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": "Something bad happened"})
		return
	}

	userResponse := &dto.UserResponse{
		ID:        newUser.ID,
		Name:      newUser.Name,
		Email:     newUser.Email,
		Photo:     newUser.Photo,
		Role:      dto.UserRole(newUser.Role),
		Provider:  newUser.Provider,
		CreatedAt: newUser.CreatedAt,
		UpdatedAt: newUser.UpdatedAt,
	}
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": gin.H{"user": userResponse}})
}

// SignInUser godoc
//
//	@Summary		SignInUser
//	@Description	SignInUser
//	@Tags			auth
//	@Param			payload	body		dto.SignInInput	true	"Login payload"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}		string
//	@Router			/auth/login [post]
func (ac *AuthController) SignInUser(ctx *gin.Context) {
	var payload *dto.SignInInput

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	user, err2 := ac.service.FindUserByEmail(payload.Email)
	if err2 != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid email or Password"})
		return
	}

	if err := utils.VerifyPassword(user.Password, payload.Password); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid email or Password"})
		return
	}

	config, _ := initializers.LoadConfig(".")

	// Generate Tokens
	accessToken, err := utils.CreateToken(config.AccessTokenExpiresIn, user.ID, config.AccessTokenPrivateKey)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	refreshToken, err := utils.CreateToken(config.RefreshTokenExpiresIn, user.ID, config.RefreshTokenPrivateKey)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.SetCookie("access_token", accessToken, config.AccessTokenMaxAge*60, "/", "localhost", false, true)
	ctx.SetCookie("refresh_token", refreshToken, config.RefreshTokenMaxAge*60, "/", "localhost", false, true)
	ctx.SetCookie("logged_in", "true", config.AccessTokenMaxAge*60, "/", "localhost", false, false)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "access_token": accessToken, "refresh_token": refreshToken})
}

// RefreshAccessToken godoc
//
//	@Summary		RefreshAccessToken
//	@Description	RefreshAccessToken
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	string
//	@Router			/auth/refresh [get]
func (ac *AuthController) RefreshAccessToken(ctx *gin.Context) {
	message := "could not refresh access token"

	cookie, err := ctx.Cookie("refresh_token")

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": message})
		return
	}

	config, _ := initializers.LoadConfig(".")

	sub, err := utils.ValidateToken(cookie, config.RefreshTokenPublicKey)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	user, result := ac.service.FindUserById(sub)

	if result != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": "the user belonging to this token no logger exists"})
		return
	}

	access_token, err := utils.CreateToken(config.AccessTokenExpiresIn, user.ID, config.AccessTokenPrivateKey)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.SetCookie("access_token", access_token, config.AccessTokenMaxAge*60, "/", "localhost", false, true)
	ctx.SetCookie("logged_in", "true", config.AccessTokenMaxAge*60, "/", "localhost", false, false)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "access_token": access_token})
}

// LogoutUser godoc
//
//	@Summary		LogoutUser
//	@Description	LogoutUser
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	string
//	@Router			/auth/logout [get]
func (ac *AuthController) LogoutUser(ctx *gin.Context) {
	ctx.SetCookie("access_token", "", -1, "/", "localhost", false, true)
	ctx.SetCookie("refresh_token", "", -1, "/", "localhost", false, true)
	ctx.SetCookie("logged_in", "", -1, "/", "localhost", false, false)

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}
