package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/VuThanhThien/golang-gorm-postgres/authentication/internal/api/services"
	"github.com/VuThanhThien/golang-gorm-postgres/authentication/internal/middleware"
	"github.com/VuThanhThien/golang-gorm-postgres/authentication/internal/models"
	"github.com/VuThanhThien/golang-gorm-postgres/authentication/internal/models/dto"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	DB      *gorm.DB
	service services.UserServiceInterface
}

func NewUserController(DB *gorm.DB, service services.UserServiceInterface) UserController {
	return UserController{DB, service}
}

// ListUsers godoc
//
//		@Summary		ListUsers
//		@Description	ListUsers
//		@Tags			users
//		@Accept			json
//		@Produce		json
//		@Param			payload	query		dto.ListUserDto			false	"ListOrders payload"
//		@Param 			_ 		query 		dto.PaginationDto 		false 	"PaginationDto"
//		@Success		200	{object}		dto.UserListResponse
//	 	@Security		Bearer
//		@Router			/users/list [get]
func (uc *UserController) ListUsers(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	var payload dto.ListUserDto
	var pagination dto.PaginationDto

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}
	pagination.Page = &page
	pagination.PageSize = &pageSize
	payload.Email = c.Query("email")
	payload.Name = c.Query("name")
	role := c.Query("role")
	if role != "" {
		payload.Role = dto.UserRole(role)
	}
	payload.Provider = c.Query("provider")

	users, total, err := uc.service.ListUsers(payload, pagination)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := dto.ToUserListResponse(users)
	response.Total = total

	c.JSON(http.StatusOK, response)
}

// GetMe godoc
//
//		@Summary		GetMe
//		@Description	GetMe
//		@Tags			users
//		@Accept			json
//		@Produce		json
//		@Success		200	{object}	dto.UserResponse
//		@Failure 		500 {string} 	string 				"an error occurred during the modification"
//	 	@Security		Bearer
//		@Router			/users/me [get]
func (uc *UserController) GetMe(ctx *gin.Context) {
	currentUser := ctx.MustGet(middleware.CURRENT_USER).(models.User)
	userResponse := &dto.UserResponse{
		ID:        currentUser.ID,
		Name:      currentUser.Name,
		Email:     currentUser.Email,
		Photo:     currentUser.Photo,
		Role:      dto.UserRole(currentUser.Role),
		Provider:  currentUser.Provider,
		CreatedAt: currentUser.CreatedAt,
		UpdatedAt: currentUser.UpdatedAt,
	}
	fmt.Printf("%v", userResponse)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"user": userResponse}})
}

// GetUser godoc
//
//		@Summary		GetUser
//		@Description	GetUser
//		@Tags			users
//		@Accept			json
//		@Produce		json
//		@Param			id	path		string	true	"User ID"
//		@Success		200	{object}	dto.UserResponse
//		@Failure 		500 {string} 	string 				"an error occurred during the modification"
//	 	@Security		Bearer
//		@Router			/users/{id} [get]
func (uc *UserController) GetUser(ctx *gin.Context) {
	var readUserRequest dto.ReadUserRequest
	if err := ctx.ShouldBindUri(&readUserRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := uc.service.ReadUser(uint(readUserRequest.ID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, dto.ToUserResponse(user))
}
