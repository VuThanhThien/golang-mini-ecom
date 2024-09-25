package services

import (
	"errors"
	"fmt"

	"github.com/VuThanhThien/golang-gorm-postgres/merchant/internal/api/repositories"
	"github.com/VuThanhThien/golang-gorm-postgres/merchant/internal/models"
	"github.com/VuThanhThien/golang-gorm-postgres/merchant/internal/models/dto"
	"gorm.io/gorm"
)

type ICategoryService interface {
	GetCategoryByID(id uint) (*models.Category, error)
	GetCategoryByCategoryID(categoryID uint) (*models.Category, error)
	CreateCategory(categoryDTO *dto.CategoryDTO) (*dto.CategoryDTO, error)
	DeleteCategory(id uint) error
	ListCategory(dto dto.ListCategoryDto, pagination dto.PaginationDto) ([]models.Category, int, error)
}

type CategoryService struct {
	categoryRepository *repositories.CategoryRepository
}

func NewCategoryService(categoryRepository *repositories.CategoryRepository) *CategoryService {
	return &CategoryService{categoryRepository: categoryRepository}
}

func (s *CategoryService) GetCategoryByCategoryID(categoryID uint) (*models.Category, error) {
	return s.categoryRepository.GetByCategoryID(categoryID)
}

func (s *CategoryService) GetCategoryByID(id uint) (*models.Category, error) {
	return s.categoryRepository.GetByID(id)
}

func (s *CategoryService) CreateCategory(categoryDTO *dto.CategoryDTO) (*dto.CategoryDTO, error) {
	category := &models.Category{
		Name:        categoryDTO.Name,
		Description: categoryDTO.Description,
	}

	// Only set ParentID if it's provided and not 0
	if categoryDTO.ParentID != nil && *categoryDTO.ParentID != 0 {
		// Check if parent category exists
		parentCategory := &models.Category{}
		if err := s.categoryRepository.GetDB().First(parentCategory, *categoryDTO.ParentID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, fmt.Errorf("parent category with ID %d not found", *categoryDTO.ParentID)
			}
			return nil, err
		}
		category.ParentID = *categoryDTO.ParentID
	}

	if err := s.categoryRepository.GetDB().Create(category).Error; err != nil {
		return nil, err
	}

	// Convert the created category back to DTO
	createdCategoryDTO := &dto.CategoryDTO{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
		ParentID:    &category.ParentID,
	}

	return createdCategoryDTO, nil
}

func (s *CategoryService) DeleteCategory(id uint) error {
	return s.categoryRepository.Delete(id)
}

func (s *CategoryService) ListCategory(dto dto.ListCategoryDto, pagination dto.PaginationDto) ([]models.Category, int, error) {
	return s.categoryRepository.List(dto, pagination)
}
