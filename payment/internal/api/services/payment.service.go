package services

import (
	"time"

	"github.com/VuThanhThien/golang-gorm-postgres/payment/internal/api/repositories"
	"github.com/VuThanhThien/golang-gorm-postgres/payment/internal/models"
	"github.com/VuThanhThien/golang-gorm-postgres/payment/internal/models/dto"
	"github.com/VuThanhThien/golang-gorm-postgres/payment/pkg/rabbitmq"
	"github.com/rs/zerolog/log"
	"golang.org/x/exp/rand"
)

type PaymentServiceInterface interface {
	CreatePayment(dto dto.CreatePaymentDto) (*models.Payment, error)
	ListPayments(dto dto.FilterPaymentDto, page, pageSize int) (*dto.PaginationResult, error)
	ReadPayment(id uint) (*models.Payment, error)
	ReadByOrderId(orderId uint) ([]models.Payment, error)
}

type PaymentService struct {
	repo                           *repositories.PaymentRepository
	paymentOrderCompletedPublisher rabbitmq.IPublisher
}

func NewPaymentService(repo *repositories.PaymentRepository, paymentOrderCompletedPublisher rabbitmq.IPublisher) *PaymentService {
	return &PaymentService{repo: repo, paymentOrderCompletedPublisher: paymentOrderCompletedPublisher}
}

func (s *PaymentService) CreatePayment(dto dto.CreatePaymentDto) (*models.Payment, error) {
	// Generate a random payment ID
	paymentID := generateRandomPaymentID()

	payment := models.Payment{
		PaymentID:     paymentID,
		OrderID:       dto.OrderID,
		Amount:        dto.Amount,
		Currency:      string(dto.Currency),
		Method:        string(dto.Method),
		Status:        string(dto.Status),
		TransactionID: dto.TransactionID,
		PaidAt:        time.Now(),
	}

	//? Mock payment status by amount here
	if payment.Amount > 100 {
		payment.Status = "FAILED"
	} else {
		payment.Status = "SUCCESS"
	}

	err := s.repo.Create(&payment)
	if err := s.paymentOrderCompletedPublisher.PublishMessage(payment); err != nil {
		log.Error().Err(err).Msg("Publish message for payment order completed error")
		return nil, err
	}
	return &payment, err
}

func (s *PaymentService) ListPayments(dto dto.FilterPaymentDto, page, pageSize int) (*dto.PaginationResult, error) {
	return s.repo.ListPaymentsWithFilter(dto, page, pageSize)
}

func (s *PaymentService) ReadPayment(id uint) (*models.Payment, error) {
	payment, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return payment, nil
}

func (s *PaymentService) ReadByOrderId(orderId uint) ([]models.Payment, error) {
	payment, err := s.repo.GetByOrderId(orderId)
	if err != nil {
		return nil, err
	}
	return payment, nil
}

func generateRandomPaymentID() string {
	const prefix = "payment_"
	const length = 16
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	randomString := make([]byte, length)
	for i := range randomString {
		randomString[i] = charset[rand.Intn(len(charset))]
	}

	return prefix + string(randomString)
}
