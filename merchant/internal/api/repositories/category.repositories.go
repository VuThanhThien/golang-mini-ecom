package repositories

import (
	"github.com/VuThanhThien/golang-gorm-postgres/merchant/internal/models"
	"github.com/VuThanhThien/golang-gorm-postgres/merchant/internal/models/dto"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CategoryRepository struct {
	BaseRepository[models.Category]
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{
		BaseRepository: NewBaseRepository[models.Category](db),
	}
}

func (r *CategoryRepository) GetByName(name string) (*models.Category, error) {
	var category models.Category
	err := r.GetDB().Where("name = ?", name).First(&category).Error
	return &category, err
}

func (r *CategoryRepository) GetByCategoryID(categoryID uint) (*models.Category, error) {
	var category models.Category
	err := r.GetDB().Where("category_id = ?", categoryID).First(&category).Error
	return &category, err
}

func (r *CategoryRepository) List(dto dto.ListCategoryDto, pagination dto.PaginationDto) ([]models.Category, int, error) {

	var categories []models.Category
	var total int64

	offset := (*pagination.Page - 1) * *pagination.PageSize

	clauses := make([]clause.Expression, 0)
	if dto.ID != 0 {
		clauses = append(clauses, clause.Eq{Column: "id", Value: dto.ID})
	}
	if dto.Name != "" {
		clauses = append(clauses, clause.Like{Column: "name", Value: "%" + dto.Name + "%"})
	}
	if dto.ParentID != 0 {
		clauses = append(clauses, clause.Eq{Column: "parent_id", Value: dto.ParentID})
	}

	err := r.GetDB().Model(&models.Category{}).Clauses(clauses...).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.GetDB().
		Model(&models.Category{}).
		Clauses(clauses...).
		Offset(offset).
		Limit(*pagination.PageSize).
		Find(&categories).
		Error
	if err != nil {
		return nil, 0, err
	}

	return categories, int(total), nil
}
