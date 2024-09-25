package repositories

import (
	"github.com/VuThanhThien/golang-gorm-postgres/merchant/internal/models"
	"github.com/VuThanhThien/golang-gorm-postgres/merchant/internal/models/dto"
	"gorm.io/gorm"
)

// ProductRepository struct
type ProductRepository struct {
	BaseRepository[models.Product]
}

// NewProductRepository creates a new product repository
func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		BaseRepository: NewBaseRepository[models.Product](db),
	}
}

// GetByName retrieves a product by its name
func (r *ProductRepository) GetByName(name string) (*models.Product, error) {
	var product models.Product
	err := r.GetDB().Where("name = ?", name).First(&product).Error
	return &product, err
}

// GetAllByCategory retrieves all products in a specific category
func (r *ProductRepository) GetAllByCategory(category string) ([]models.Product, error) {
	var products []models.Product
	err := r.GetDB().Where("category = ?", category).Find(&products).Error
	return products, err
}

// UpdateStock updates the stock quantity of a product
func (r *ProductRepository) UpdateStock(id uint, quantity int) error {
	return r.GetDB().Model(&models.Product{}).Where("id = ?", id).Update("stock", quantity).Error
}

// FilterProductsWithPagination filters products based on given criteria with pagination
func (r *ProductRepository) FilterProductsWithPagination(filter dto.FilterOptions, page, pageSize int) (*dto.PaginationResult, error) {
	var products []models.Product
	var totalItems int64

	query := r.GetDB().Model(&models.Product{}).Preload("Variants")

	// Apply filters
	if filter.Name != "" {
		query = query.Where("name LIKE ?", "%"+filter.Name+"%")
	}
	if filter.CategoryId != "" {
		query = query.Where("category_id = ?", filter.CategoryId)
	}
	if filter.MinPrice > 0 {
		query = query.Where("price >= ?", filter.MinPrice)
	}
	if filter.MaxPrice > 0 {
		query = query.Where("price <= ?", filter.MaxPrice)
	}

	// Count total items
	if err := query.Count(&totalItems).Error; err != nil {
		return nil, err
	}

	// Calculate pagination
	offset := (page - 1) * pageSize
	totalPages := int((totalItems + int64(pageSize) - 1) / int64(pageSize))

	// Retrieve paginated data
	err := query.Offset(offset).Limit(pageSize).Find(&products).Error
	if err != nil {
		return nil, err
	}

	return &dto.PaginationResult{
		Data:       products,
		TotalItems: totalItems,
		TotalPages: totalPages,
		Page:       page,
		PageSize:   pageSize,
	}, nil
}

// GetByProductID retrieves a product by its product ID
func (r *ProductRepository) GetByProductID(productID uint) (*models.Product, error) {
	var product models.Product
	err := r.GetDB().Where("product_id = ?", productID).First(&product).Error
	return &product, err
}

// CreateProduct creates a new product
func (r *ProductRepository) CreateProduct(product *models.Product) (*models.Product, error) {
	return product, r.GetDB().Create(product).Error
}
