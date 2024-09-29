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

func (r *OrderRepository) GetByID(id uint) (*models.Order, error) {
	var order models.Order
	err := r.GetDB().Preload("Items").Where("id = ?", id).First(&order).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *OrderRepository) GetByOrderId(orderId string) (*models.Order, error) {
	var order models.Order
	err := r.GetDB().Preload("Items").Where("order_id = ?", orderId).First(&order).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *OrderRepository) GetOrders(orderDto dto.GetOrderRequestDto, page, pageSize int) (*dto.PaginationResult, error) {
	var orders []models.Order
	query := r.GetDB().Model(&models.Order{}).Preload("Items")

	if orderDto.UserID != 0 {
		query = query.Where("user_id = ?", orderDto.UserID)
	}

	if orderDto.MerchantID != 0 {
		query = query.Where("merchant_id = ?", orderDto.MerchantID)
	}

	if orderDto.Status != "" {
		query = query.Where("status = ?", orderDto.Status)
	}

	if orderDto.PaymentID != 0 {
		query = query.Where("payment_id = ?", orderDto.PaymentID)
	}

	if orderDto.OrderID != 0 {
		query = query.Where("order_id = ?", orderDto.OrderID)
	}

	var total int64
	err := query.Count(&total).Error
	if err != nil {
		return nil, err
	}

	offset := (page - 1) * pageSize
	totalPages := int((total + int64(pageSize) - 1) / int64(pageSize))

	err = query.Limit(pageSize).Offset(offset).Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return &dto.PaginationResult{
		Data:       orders,
		TotalItems: total,
		TotalPages: totalPages,
		Page:       page,
		PageSize:   pageSize,
	}, nil
}

func (r *OrderRepository) CreateWithTx(tx *gorm.DB, order *models.Order) error {
	return tx.Session(&gorm.Session{FullSaveAssociations: true}).Create(order).Error
}

func (r *OrderRepository) UpdateWithTx(tx *gorm.DB, merchant *models.Order) error {
	return tx.Model(merchant).Updates(merchant).Error
}
