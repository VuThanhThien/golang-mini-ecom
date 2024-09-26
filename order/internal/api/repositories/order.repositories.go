package repositories

import (
	"github.com/VuThanhThien/golang-gorm-postgres/order/internal/models"
	"github.com/VuThanhThien/golang-gorm-postgres/order/internal/models/dto"
	"gorm.io/gorm"
)

type OrderRepository struct {
	BaseRepository[models.Order]
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{BaseRepository: NewBaseRepository[models.Order](db)}
}

func (r *OrderRepository) GetByOrderId(orderId string) (*models.Order, error) {
	var order models.Order
	err := r.GetDB().Where("order_id = ?", orderId).First(&order).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *OrderRepository) GetOrders(dto dto.GetOrderRequestDto) ([]models.Order, error) {
	var orders []models.Order
	query := r.GetDB()

	if dto.UserID != 0 {
		query = query.Where("user_id = ?", dto.UserID)
	}

	if dto.MerchantID != 0 {
		query = query.Where("merchant_id = ?", dto.MerchantID)
	}

	if dto.Status != "" {
		query = query.Where("status = ?", dto.Status)
	}

	if dto.PaymentID != 0 {
		query = query.Where("payment_id = ?", dto.PaymentID)
	}

	if dto.OrderID != 0 {
		query = query.Where("order_id = ?", dto.OrderID)
	}

	err := query.Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *OrderRepository) CreateWithTx(tx *gorm.DB, order *models.Order) error {
	return tx.Session(&gorm.Session{FullSaveAssociations: true}).Create(order).Error
}

func (r *OrderRepository) UpdateWithTx(tx *gorm.DB, merchant *models.Order) error {
	return tx.Model(merchant).Updates(merchant).Error
}
