package repositories

import (
	"github.com/VuThanhThien/golang-gorm-postgres/merchant/internal/models"
	"gorm.io/gorm"
)

type VariantRepository struct {
	BaseRepository[models.Variant]
}

func NewVariantRepository(db *gorm.DB) *VariantRepository {
	return &VariantRepository{
		BaseRepository: NewBaseRepository[models.Variant](db),
	}
}

func (r *VariantRepository) GetByProductID(productID uint) (*models.Variant, error) {
	var variant models.Variant
	err := r.GetDB().Where("product_id = ?", productID).First(&variant).Error
	return &variant, err
}

func (r *VariantRepository) GetByVariantName(variantName string) (*models.Variant, error) {
	var variant models.Variant
	err := r.GetDB().Where("variant_name = ?", variantName).First(&variant).Error
	return &variant, err
}
