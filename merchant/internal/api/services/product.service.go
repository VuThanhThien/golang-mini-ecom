package services

import (
	"github.com/VuThanhThien/golang-gorm-postgres/merchant/internal/api/repositories"
	"github.com/VuThanhThien/golang-gorm-postgres/merchant/internal/models"
	"github.com/VuThanhThien/golang-gorm-postgres/merchant/internal/models/dto"
)

type IProductService interface {
	GetProductByID(id uint) (*models.Product, error)
	GetProductByProductID(productID uint) (*models.Product, error)
	CreateProduct(dto *dto.CreateProductInput) (*models.Product, error)
	DeleteProduct(id uint) error
	FilterProductsWithPagination(filter dto.FilterOptions, page, pageSize int) (*dto.PaginationResult, error)
	UpdateStock(id uint, quantity int) error
	GetAllByCategory(category string) ([]models.Product, error)
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

func (s *ProductService) CreateProduct(dto *dto.CreateProductInput) (*models.Product, error) {
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

func (s *ProductService) FilterProductsWithPagination(filter dto.FilterOptions, page, pageSize int) (*dto.PaginationResult, error) {
	return s.productRepository.FilterProductsWithPagination(filter, page, pageSize)
}

func (s *ProductService) UpdateStock(id uint, quantity int) error {
	return s.productRepository.UpdateStock(id, quantity)
}

func (s *ProductService) GetAllByCategory(category string) ([]models.Product, error) {
	return s.productRepository.GetAllByCategory(category)
}
