package services

import (
	"github.com/VuThanhThien/golang-gorm-postgres/merchant/internal/api/repositories"
	"github.com/VuThanhThien/golang-gorm-postgres/merchant/internal/models"
	"github.com/VuThanhThien/golang-gorm-postgres/merchant/internal/models/dto"
	"github.com/VuThanhThien/golang-gorm-postgres/merchant/pkg/rabbitmq"
	"gorm.io/gorm"
)

type IMerchantService interface {
	GetMerchantByID(id uint) (*models.Merchant, error)
	GetMerchantByMerchantID(merchantID string) (*models.Merchant, error)
	CreateMerchantWithTx(dto *dto.CreateMerchantDTO) (*models.Merchant, error)
	UpdateMerchantWithTx(dto *dto.UpdateMerchantDTO) (*models.Merchant, error)
}

type MerchantService struct {
	repo                    *repositories.MerchantRepository
	CreateMerchantPublisher rabbitmq.IPublisher
}

func NewMerchantService(repo *repositories.MerchantRepository, createMerchantPublisher rabbitmq.IPublisher) *MerchantService {
	return &MerchantService{repo: repo, CreateMerchantPublisher: createMerchantPublisher}
}

func (s *MerchantService) GetMerchantByID(id uint) (*models.Merchant, error) {
	merchant, err := s.repo.GetByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return merchant, nil
}

func (s *MerchantService) GetMerchantByMerchantID(merchantID string) (*models.Merchant, error) {
	return s.repo.GetByMerchantID(merchantID)
}

func (s *MerchantService) CreateMerchantWithTx(dto *dto.CreateMerchantDTO) (*models.Merchant, error) {
	merchant := &models.Merchant{
		Name:        dto.Name,
		MerchantID:  dto.MerchantID,
		Description: dto.Description,
		UserID:      dto.UserID,
	}

	tx := s.repo.GetDB().Begin()
	if err := s.repo.CreateWithTx(tx, merchant); err != nil {
		tx.Rollback()
		return nil, err
	}
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	if err := s.CreateMerchantPublisher.PublishMessage(merchant); err != nil {
		return nil, err
	}
	return merchant, nil
}

func (s *MerchantService) UpdateMerchantWithTx(dto *dto.UpdateMerchantDTO) (*models.Merchant, error) {
	merchant := &models.Merchant{
		Name:        dto.Name,
		MerchantID:  dto.MerchantID,
		Description: dto.Description,
	}

	tx := s.repo.GetDB().Begin()
	if err := s.repo.UpdateWithTx(tx, merchant); err != nil {
		tx.Rollback()
		return nil, err
	}
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}
	return merchant, nil
}
