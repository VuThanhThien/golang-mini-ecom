package services

import (
	"github.com/VuThanhThien/golang-gorm-postgres/merchant/internal/api/repositories"
	"github.com/VuThanhThien/golang-gorm-postgres/merchant/internal/models"
	"github.com/VuThanhThien/golang-gorm-postgres/merchant/internal/models/dto"
)

type IProductImageService interface {
	GetProductImageByID(id uint) (*models.ProductImage, error)
	GetProductImageByProductID(productID uint) (*models.ProductImage, error)
	CreateProductImage(dto *dto.ProductImageDTO) (*models.ProductImage, error)
	DeleteProductImage(id uint) error
	UpdateProductImage(dto *dto.ProductImageDTO) (*models.ProductImage, error)
}

type ProductImageService struct {
	productImageRepository *repositories.ProductImageRepository
}

func NewProductImageService(productImageRepository *repositories.ProductImageRepository) *ProductImageService {
	return &ProductImageService{productImageRepository: productImageRepository}
}

func (s *ProductImageService) GetProductImageByID(id uint) (*models.ProductImage, error) {
	return s.productImageRepository.GetByID(id)
}

func (s *ProductImageService) GetProductImageByProductID(productID uint) (*models.ProductImage, error) {
	return s.productImageRepository.GetByProductID(productID)
}

func (s *ProductImageService) CreateProductImage(dto *dto.ProductImageDTO) (*models.ProductImage, error) {
	return s.productImageRepository.Create(dto)
}

func (s *ProductImageService) DeleteProductImage(id uint) error {
	return s.productImageRepository.Delete(id)
}

func (s *ProductImageService) UpdateProductImage(dto *dto.ProductImageDTO) (*models.ProductImage, error) {
	return s.productImageRepository.Update(dto)
}
