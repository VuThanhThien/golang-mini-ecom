package services

import (
	"strings"

	"github.com/VuThanhThien/golang-gorm-postgres/internal/api/repositories"
	"github.com/VuThanhThien/golang-gorm-postgres/internal/models"
	_ "gorm.io/gorm"
)

type AuthServiceInterface interface {
	SignUpUser(order *models.User) error
	FindUserByEmail(email string) (*models.User, error)
}

type AuthService struct {
	repo *repositories.AuthRepository
}

func NewAuthService(repo *repositories.AuthRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) SignUpUser(user *models.User) error {

	tx := s.repo.GetDB().Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := s.repo.CreateWithTx(tx, user); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (s *AuthService) FindUserByEmail(email string) (*models.User, error) {
	var user models.User

	err := s.repo.GetDB().First(&user, "email = ?", strings.ToLower(email)).Error

	return &user, err
}
