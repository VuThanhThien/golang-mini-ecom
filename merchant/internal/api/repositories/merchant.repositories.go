package repositories

import (
	"github.com/VuThanhThien/golang-gorm-postgres/merchant/internal/models"
	"gorm.io/gorm"
)

type MerchantRepository struct {
	db *gorm.DB
}

func (r *MerchantRepository) GetDB() *gorm.DB {
	return r.db
}

func NewMerchantRepository(db *gorm.DB) *MerchantRepository {
	return &MerchantRepository{db: db}
}

func (r *MerchantRepository) Create(merchant *models.Merchant) error {
	return r.db.Create(merchant).Error
}

func (r *MerchantRepository) GetByID(id string) (*models.Merchant, error) {
	var merchant models.Merchant
	err := r.db.Where("id::text = ?", id).First(&merchant).Error
	return &merchant, err
}

func (r *MerchantRepository) GetByMerchantID(merchantID string) (*models.Merchant, error) {
	var merchant models.Merchant
	err := r.db.Where("merchant_id = ?", merchantID).First(&merchant).Error
	return &merchant, err
}

func (r *MerchantRepository) CreateWithTx(tx *gorm.DB, merchant *models.Merchant) error {
	return tx.Session(&gorm.Session{FullSaveAssociations: true}).Create(merchant).Error
}
