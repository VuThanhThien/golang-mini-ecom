package repositories

import (
	"github.com/VuThanhThien/golang-gorm-postgres/order/internal/models"
	"gorm.io/gorm"
)

type ItemRepository struct {
	BaseRepository[models.Item]
}

func NewItemRepository(db *gorm.DB) *ItemRepository {
	return &ItemRepository{BaseRepository: NewBaseRepository[models.Item](db)}
}

func (r *ItemRepository) CreateWithTx(tx *gorm.DB, item *models.Item) error {
	return tx.Session(&gorm.Session{FullSaveAssociations: true}).Create(item).Error
}

func (r *ItemRepository) GetByOrderId(orderId string) ([]models.Item, error) {
	var items []models.Item
	err := r.GetDB().Where("order_id = ?", orderId).Find(&items).Error
	if err != nil {
		return nil, err
	}
	return items, nil
}
