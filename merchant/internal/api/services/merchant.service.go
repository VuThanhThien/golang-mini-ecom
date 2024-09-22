package services

import (
	"github.com/VuThanhThien/golang-gorm-postgres/merchant/internal/api/repositories"
	"github.com/VuThanhThien/golang-gorm-postgres/merchant/internal/models"
	"github.com/VuThanhThien/golang-gorm-postgres/merchant/internal/models/dto"
)

type MerchantServiceInterface interface {
	GetMerchantByID(id string) (*models.Merchant, error)
	GetMerchantByMerchantID(merchantID string) (*models.Merchant, error)
	CreateMerchantWithTx(dto *dto.CreateMerchantDTO) (*models.Merchant, error)
}

type MerchantService struct {
	repo *repositories.MerchantRepository
}

func NewMerchantService(repo *repositories.MerchantRepository) *MerchantService {
	return &MerchantService{repo: repo}
}

func (s *MerchantService) GetMerchantByID(id string) (*models.Merchant, error) {
	return s.repo.GetByID(id)
}

func (s *MerchantService) GetMerchantByMerchantID(merchantID string) (*models.Merchant, error) {
	return s.repo.GetByMerchantID(merchantID)
}

func (s *MerchantService) CreateMerchantWithTx(dto *dto.CreateMerchantDTO) (*models.Merchant, error) {
	merchant := &models.Merchant{
		Name:        dto.Name,
		MerchantId:  dto.MerchantId,
		Description: dto.Description,
	}

	tx := s.repo.GetDB().Begin()
	if err := s.repo.CreateWithTx(tx, merchant); err != nil {
		tx.Rollback()
		return nil, err
	}
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}
	return merchant, nil
}
