package services

import (
	"github.com/VuThanhThien/golang-gorm-postgres/merchant/internal/api/repositories"
	"github.com/VuThanhThien/golang-gorm-postgres/merchant/internal/models"
	"github.com/VuThanhThien/golang-gorm-postgres/merchant/internal/models/dto"
)

type IVariantService interface {
	Create(variantDTO dto.VariantDTO) (*models.Variant, error)
	GetByProductID(productID uint) (*models.Variant, error)
	GetByVariantName(variantName string) (*models.Variant, error)
	Delete(id uint) error
	GetById(id uint) (*models.Variant, error)
}

type VariantService struct {
	variantRepository *repositories.VariantRepository
}

func NewVariantService(variantRepository *repositories.VariantRepository) *VariantService {
	return &VariantService{
		variantRepository: variantRepository,
	}
}

func (s *VariantService) Create(variantDTO dto.VariantDTO) (*models.Variant, error) {

	variant := &models.Variant{
		ProductID:   variantDTO.ProductID,
		VariantName: variantDTO.VariantName,
		Description: variantDTO.Description,
		Price:       variantDTO.Price,
	}

	err := s.variantRepository.Create(variant)
	return variant, err
}

func (s *VariantService) GetByProductID(productID uint) (*models.Variant, error) {
	return s.variantRepository.GetByProductID(productID)
}

func (s *VariantService) GetByVariantName(variantName string) (*models.Variant, error) {
	return s.variantRepository.GetByVariantName(variantName)
}

func (s *VariantService) Delete(id uint) error {
	return s.variantRepository.Delete(id)
}

func (s *VariantService) GetById(id uint) (*models.Variant, error) {
	return s.variantRepository.GetByID(id)
}
