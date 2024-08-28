package controllers

import (
	"fmt"
	"net/http"

	"github.com/VuThanhThien/golang-gorm-postgres/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

func NewUserController(DB *gorm.DB) UserController {
	return UserController{DB}
}

// GetMe godoc
//
//		@Summary		GetMe
//		@Description	GetMe
//		@Tags			users
//		@Accept			json
//		@Produce		json
//		@Success		200	{object}	models.UserResponse
//	 	@Security		Bearer
//		@Router			/users/me [get]
func (uc *UserController) GetMe(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)

	userResponse := &models.UserResponse{
		ID:        currentUser.ID,
		Name:      currentUser.Name,
		Email:     currentUser.Email,
		Photo:     currentUser.Photo,
		Role:      currentUser.Role,
		Provider:  currentUser.Provider,
		CreatedAt: currentUser.CreatedAt,
		UpdatedAt: currentUser.UpdatedAt,
	}
	fmt.Printf("%v", userResponse)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"user": userResponse}})
}
