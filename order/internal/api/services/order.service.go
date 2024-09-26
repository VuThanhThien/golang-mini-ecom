package services

import (
	"github.com/VuThanhThien/golang-gorm-postgres/order/internal/api/repositories"
	"github.com/VuThanhThien/golang-gorm-postgres/order/internal/models"
	"github.com/VuThanhThien/golang-gorm-postgres/order/internal/models/dto"
	"github.com/VuThanhThien/golang-gorm-postgres/order/pkg/rabbitmq"
	"gorm.io/gorm"
)

type IOrderService interface {
	GetOrder(id uint) (*models.Order, error)
	GetOrderByOrderID(orderID string) (*models.Order, error)
	CreateOrder(createOrderDto dto.CreateOrderRequestDto) (*models.Order, error)
}

type OrderService struct {
	orderRepository      *repositories.OrderRepository
	orderItemRepository  *repositories.ItemRepository
	CreateOrderPublisher rabbitmq.IPublisher
}

func NewOrderService(orderRepo *repositories.OrderRepository, orderItemRepo *repositories.ItemRepository, createOrderPublisher rabbitmq.IPublisher) *OrderService {
	return &OrderService{orderRepository: orderRepo, orderItemRepository: orderItemRepo, CreateOrderPublisher: createOrderPublisher}
}

func (s *OrderService) GetOrder(id uint) (*models.Order, error) {
	merchant, err := s.orderRepository.GetByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return merchant, nil
}

func (s *OrderService) GetOrderByOrderID(orderID string) (*models.Order, error) {
	return s.orderRepository.GetByOrderId(orderID)
}

func (s *OrderService) CreateOrder(createOrderDto dto.CreateOrderRequestDto) (*models.Order, error) {
	order := &models.Order{
		OrderID:     createOrderDto.OrderID,
		UserID:      createOrderDto.UserID,
		TotalAmount: createOrderDto.TotalAmount,
		Status:      "PENDING",
	}

	err := s.orderRepository.GetDB().Transaction(func(tx *gorm.DB) error {
		return s.orderRepository.CreateWithTx(tx, order)
	})
	if err != nil {
		return nil, err
	}

	// Save items to the order
	for _, itemDto := range createOrderDto.Items {
		item := &models.Item{
			OrderID:    order.OrderID,
			ProductID:  itemDto.ProductID,
			VariantID:  itemDto.VariantID,
			Name:       itemDto.Name,
			Quantity:   itemDto.Quantity,
			Price:      itemDto.Price,
			TotalPrice: itemDto.TotalPrice,
		}
		if err := s.orderItemRepository.Create(item); err != nil {
			return nil, err
		}
	}

	if err := s.CreateOrderPublisher.PublishMessage(order); err != nil {
		return nil, err
	}

	return order, nil
}

func (s *OrderService) UpdateOrder(dto dto.UpdateOrderRequestDto) (*models.Order, error) {
	order, err := s.orderRepository.GetByID(dto.ID)
	if err != nil {
		return nil, err
	}
	order.Status = string(dto.Status)
	return order, s.orderRepository.Update(order)
}
