package dto

import (
	"time"

	"github.com/VuThanhThien/golang-gorm-postgres/authentication/internal/models"
)

type SignUpInput struct {
	Name            string `json:"name" binding:"required" example:"admin"`
	Email           string `json:"email" binding:"required" example:"admin@gmail.com"`
	Password        string `json:"password" binding:"required,min=8" example:"123456@Abc"`
	PasswordConfirm string `json:"passwordConfirm" binding:"required" example:"123456@Abc"`
	Photo           string `json:"photo" binding:"required"`
}

type SignInInput struct {
	Email    string `json:"email"  binding:"required" example:"admin@gmail.com"`
	Password string `json:"password"  binding:"required" example:"123456@Abc"`
}

type ReadUserRequest struct {
	ID uint `uri:"id" binding:"required,min=1"`
}

type UserRole string

const (
	UserRoleAdmin UserRole = "admin"
	UserRoleUser  UserRole = "user"
)

type UserResponse struct {
	ID        uint      `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Role      UserRole  `json:"role,omitempty"`
	Photo     string    `json:"photo,omitempty"`
	Provider  string    `json:"provider"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ListUserDto struct {
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	Role     UserRole `json:"role"`
	Provider string   `json:"provider"`
}

type UserListResponse struct {
	Users []UserResponse `json:"users"`
	Total int            `json:"total"`
}

func ToUserListResponse(users []models.User) UserListResponse {
	summaries := make([]UserResponse, len(users))
	for i, user := range users {
		summaries[i] = UserResponse{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			Name:      user.Name,
			Email:     user.Email,
			Role:      UserRole(user.Role),
			Photo:     user.Photo,
			Provider:  user.Provider,
		}
	}
	return UserListResponse{
		Users: summaries,
		Total: len(summaries),
	}
}

func ToUserResponse(user *models.User) *UserResponse {
	return &UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      UserRole(user.Role),
		Photo:     user.Photo,
		Provider:  user.Provider,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
