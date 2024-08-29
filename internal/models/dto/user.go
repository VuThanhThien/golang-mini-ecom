package dto

import (
	"time"

	"github.com/VuThanhThien/golang-gorm-postgres/internal/models"
	"github.com/google/uuid"
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
	ID        uuid.UUID `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Role      string    `json:"role,omitempty"`
	Photo     string    `json:"photo,omitempty"`
	Provider  string    `json:"provider"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserListResponse struct {
	Orders []UserResponse `json:"users"`
	Total  int            `json:"total"`
}

func ToOrderListResponse(orders []models.User) UserListResponse {
	summaries := make([]UserResponse, len(orders))
	for i, order := range orders {
		summaries[i] = UserResponse{
			ID:        order.ID,
			CreatedAt: order.CreatedAt,
			UpdatedAt: order.UpdatedAt,
			Name:      order.Name,
			Email:     order.Email,
			Role:      order.Role,
			Photo:     order.Photo,
			Provider:  order.Provider,
		}
	}
	return UserListResponse{
		Orders: summaries,
		Total:  len(summaries),
	}
}
