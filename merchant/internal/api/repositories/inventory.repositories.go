package repositories

import (
	"github.com/VuThanhThien/golang-gorm-postgres/merchant/internal/models"

	"gorm.io/gorm"
)

type InventoryRepository struct {
	BaseRepository[models.Inventory]
}

func NewInventoryRepository(db *gorm.DB) *InventoryRepository {
	return &InventoryRepository{
		BaseRepository: NewBaseRepository[models.Inventory](db),
	}
}

func (r *InventoryRepository) GetByVariantID(variantID uint) (*models.Inventory, error) {
	var inventory models.Inventory
	err := r.GetDB().Where("variant_id = ?", variantID).First(&inventory).Error
	if err != nil {
		return nil, err
	}
	return &inventory, nil
}

func (r *InventoryRepository) Create(inventory *models.Inventory) (*models.Inventory, error) {
	err := r.GetDB().Create(inventory).Error
	if err != nil {
		return nil, err
	}
	err = r.GetDB().First(inventory).Error
	if err != nil {
		return nil, err
	}
	return inventory, nil
}
