package repositories

import (
	"fmt"

	"github.com/VuThanhThien/golang-gorm-postgres/payment/internal/models"
	"github.com/VuThanhThien/golang-gorm-postgres/payment/internal/models/dto"
	"gorm.io/gorm"
)

type PaymentRepository struct {
	BaseRepository[models.Payment]
}

func NewPaymentRepository(db *gorm.DB) *PaymentRepository {
	return &PaymentRepository{BaseRepository: NewBaseRepository[models.Payment](db)}
}

func (r *PaymentRepository) CreateWithTx(tx *gorm.DB, paymentDTO *dto.CreatePaymentInput) (*models.Payment, error) {
	payment := &models.Payment{
		PaymentID:     paymentDTO.PaymentID,
		Currency:      string(paymentDTO.Currency),
		Method:        string(paymentDTO.Method),
		TransactionID: paymentDTO.TransactionID,
		PaidAt:        paymentDTO.PaidAt,
		OrderID:       paymentDTO.OrderID,
		Amount:        paymentDTO.Amount,
		Status:        string(paymentDTO.Status),
	}
	err := tx.Session(&gorm.Session{FullSaveAssociations: true}).Create(payment).Error
	if err != nil {
		return nil, err
	}
	return payment, nil
}

func (r *PaymentRepository) FindByID(id uint) (*models.Payment, error) {
	payment, err := r.BaseRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return payment, nil
}

func (r *PaymentRepository) GetByOrderId(orderId uint) ([]models.Payment, error) {
	var payments []models.Payment
	err := r.GetDB().Where("order_id = ?", orderId).Find(&payments).Error
	if err != nil {
		return nil, err
	}
	return payments, nil
}
func (r *PaymentRepository) ListPaymentsWithFilter(filterDTO dto.FilterPaymentDto, page, pageSize int) (*dto.PaginationResult, error) {
	var payments []models.Payment
	query := r.GetDB().Model(&models.Payment{})

	if filterDTO.PaymentID != "" {
		query = query.Where("payment_id = ?", filterDTO.PaymentID)
	}
	if filterDTO.OrderID != 0 {
		query = query.Where("order_id = ?", filterDTO.OrderID)
	}
	if filterDTO.Currency != "" {
		query = query.Where("currency = ?", filterDTO.Currency)
	}
	if filterDTO.Method != "" {
		query = query.Where("method = ?", filterDTO.Method)
	}
	if filterDTO.Status != "" {
		query = query.Where("status = ?", filterDTO.Status)
	}
	if !filterDTO.PaidAt.IsZero() {
		query = query.Where("paid_at = ?", filterDTO.PaidAt)
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, fmt.Errorf("error counting total payments: %w", err)
	}

	offset := (page - 1) * pageSize
	totalPages := int((total + int64(pageSize) - 1) / int64(pageSize))
	err := query.Offset(offset).Limit(pageSize).Find(&payments).Error
	if err != nil {
		return nil, err
	}

	return &dto.PaginationResult{
		Data:       payments,
		TotalItems: total,
		TotalPages: totalPages,
		Page:       page,
		PageSize:   pageSize,
	}, nil
}
