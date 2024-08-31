package services

import (
	"github.com/VuThanhThien/golang-gorm-postgres/internal/api/repositories"
	"github.com/VuThanhThien/golang-gorm-postgres/internal/models"
	"github.com/VuThanhThien/golang-gorm-postgres/internal/models/dto"
)

type UserServiceInterface interface {
	ListUsers(dto dto.ListUserDto, pagination dto.PaginationDto) ([]models.User, int, error)
}

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) ListUsers(dto dto.ListUserDto, pagination dto.PaginationDto) ([]models.User, int, error) {
	return s.repo.List(dto, pagination)
}
