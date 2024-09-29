package repositories

import (
	"github.com/VuThanhThien/golang-gorm-postgres/merchant/internal/models"
	"gorm.io/gorm"
)

type MerchantRepository struct {
	BaseRepository[models.Merchant]
}

func NewMerchantRepository(db *gorm.DB) *MerchantRepository {
	return &MerchantRepository{BaseRepository: NewBaseRepository[models.Merchant](db)}
}

func (r *MerchantRepository) GetByMerchantID(merchantID string) (*models.Merchant, error) {
	var merchant models.Merchant
	err := r.GetDB().Where("merchant_id = ?", merchantID).First(&merchant).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &merchant, err
}

func (r *MerchantRepository) CreateWithTx(tx *gorm.DB, merchant *models.Merchant) error {
	return tx.Session(&gorm.Session{FullSaveAssociations: true}).Create(merchant).Error
}

func (r *MerchantRepository) UpdateWithTx(tx *gorm.DB, merchant *models.Merchant) error {
	return tx.Model(merchant).Updates(merchant).Error
}

func (r *MerchantRepository) GetByUserID(userID uint) (*models.Merchant, error) {
	var merchant models.Merchant
	err := r.GetDB().Where("user_id = ?", userID).First(&merchant).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &merchant, err
}
