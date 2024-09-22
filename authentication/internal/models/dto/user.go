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

type UserResponse struct {
	ID        string    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Role      string    `json:"role,omitempty"`
	Photo     string    `json:"photo,omitempty"`
	Provider  string    `json:"provider"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ListUserDto struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Provider string `json:"provider"`
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
			Role:      user.Role,
			Photo:     user.Photo,
			Provider:  user.Provider,
		}
	}
	return UserListResponse{
		Users: summaries,
		Total: len(summaries),
	}
}
