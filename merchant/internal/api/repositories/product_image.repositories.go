package repositories

import (
	"github.com/VuThanhThien/golang-gorm-postgres/merchant/internal/models"
	"github.com/VuThanhThien/golang-gorm-postgres/merchant/internal/models/dto"
	"gorm.io/gorm"
)

type ProductImageRepository struct {
	BaseRepository[models.ProductImage]
}

func NewProductImageRepository(db *gorm.DB) *ProductImageRepository {
	return &ProductImageRepository{
		BaseRepository: NewBaseRepository[models.ProductImage](db),
	}
}

func (r *ProductImageRepository) GetByProductID(productID uint) (*models.ProductImage, error) {
	var productImage models.ProductImage
	err := r.GetDB().Where("product_id = ?", productID).First(&productImage).Error
	return &productImage, err
}

func (r *ProductImageRepository) Create(dto *dto.ProductImageDTO) (*models.ProductImage, error) {
	productImage := models.ProductImage{
		ProductID: dto.ProductID,
		ImageURL:  dto.ImageURL,
	}
	return &productImage, r.BaseRepository.Create(&productImage)
}

func (r *ProductImageRepository) Delete(id uint) error {
	return r.BaseRepository.Delete(id)
}

func (r *ProductImageRepository) Update(dto *dto.ProductImageDTO) (*models.ProductImage, error) {
	productImage := models.ProductImage{
		ProductID: dto.ProductID,
		ImageURL:  dto.ImageURL,
	}
	return &productImage, r.BaseRepository.Update(&productImage)
}
