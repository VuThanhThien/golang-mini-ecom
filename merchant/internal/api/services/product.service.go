package services

import (
	"github.com/VuThanhThien/golang-gorm-postgres/merchant/internal/api/repositories"
	"github.com/VuThanhThien/golang-gorm-postgres/merchant/internal/models"
	"github.com/VuThanhThien/golang-gorm-postgres/merchant/internal/models/dto"
)

type IProductService interface {
	GetProductByID(id uint) (*models.Product, error)
	GetProductByProductID(productID uint) (*models.Product, error)
	CreateProduct(dto *dto.CreateProductDTO) (*models.Product, error)
	DeleteProduct(id uint) error
}

type ProductService struct {
	productRepository *repositories.ProductRepository
}

func NewProductService(productRepository *repositories.ProductRepository) *ProductService {
	return &ProductService{productRepository: productRepository}
}

func (s *ProductService) GetProductByID(id uint) (*models.Product, error) {
	return s.productRepository.GetByID(id)
}

func (s *ProductService) DeleteProduct(id uint) error {
	return s.productRepository.Delete(id)
}

func (s *ProductService) GetProductByProductID(productID uint) (*models.Product, error) {
	return s.productRepository.GetByProductID(productID)
}

func (s *ProductService) CreateProduct(dto *dto.CreateProductDTO) (*models.Product, error) {
	product := &models.Product{
		Name:        dto.Name,
		Description: dto.Description,
		Price:       dto.Price,
		SKU:         dto.SKU,
		CategoryID:  dto.CategoryID,
		MerchantID:  dto.MerchantID,
	}
	return s.productRepository.CreateProduct(product)
}
