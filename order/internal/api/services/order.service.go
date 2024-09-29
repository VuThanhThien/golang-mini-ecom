package services

import (
	"context"
	"fmt"
	"time"

	"github.com/VuThanhThien/golang-gorm-postgres/order/internal/api/repositories"
	"github.com/VuThanhThien/golang-gorm-postgres/order/internal/gateway/grpc"
	"github.com/VuThanhThien/golang-gorm-postgres/order/internal/models"
	"github.com/VuThanhThien/golang-gorm-postgres/order/internal/models/dto"
	"github.com/VuThanhThien/golang-gorm-postgres/order/pkg/rabbitmq"
	"gorm.io/gorm"
)

type IOrderService interface {
	GetOrders(dto dto.GetOrderRequestDto, page, pageSize int) (*dto.PaginationResult, error)
	GetOrder(id uint) (*models.Order, error)
	GetOrderByOrderID(orderID string) (*models.Order, error)
	CreateOrder(createOrderDto dto.CreateOrderRequestDto) (*models.Order, error)
}

type OrderService struct {
	orderRepository      *repositories.OrderRepository
	orderItemRepository  *repositories.ItemRepository
	createOrderPublisher rabbitmq.IPublisher
	inventoryGateway     grpc.IInventoryGateway
}

func NewOrderService(orderRepo *repositories.OrderRepository, orderItemRepo *repositories.ItemRepository, createOrderPublisher rabbitmq.IPublisher, inventoryGateway grpc.IInventoryGateway) *OrderService {
	return &OrderService{orderRepository: orderRepo, orderItemRepository: orderItemRepo, createOrderPublisher: createOrderPublisher, inventoryGateway: inventoryGateway}
}

func (s *OrderService) GetOrders(dto dto.GetOrderRequestDto, page, pageSize int) (*dto.PaginationResult, error) {
	return s.orderRepository.GetOrders(dto, page, pageSize)
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
	// Check inventory for each item in the order
	for _, item := range createOrderDto.Items {
		inventory, err := s.inventoryGateway.GetInventory(context.Background(), uint64(item.VariantID))
		if err != nil {
			return nil, fmt.Errorf("failed to check inventory for variant %d: %w", item.VariantID, err)
		}

		if inventory.Quantity < uint64(item.Quantity) {
			return nil, fmt.Errorf("insufficient inventory for variant %d: requested %d, available %d",
				item.VariantID, item.Quantity, inventory.Quantity)
		}
	}

	order := &models.Order{
		OrderID:     createOrderDto.OrderID,
		UserID:      createOrderDto.UserID,
		TotalAmount: createOrderDto.TotalAmount,
		Status:      "PENDING",
		PlacedAt:    time.Now(),
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
			OrderID:    order.ID,
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

	createdOrder, err := s.orderRepository.GetByID(order.ID)
	if err != nil {
		return nil, err
	}

	// TODO: Can use rabbitmq to notify other services
	// if err := s.CreateOrderPublisher.PublishMessage(createdOrder); err != nil {
	// 	log.Error().Err(err).Msg("Publish message error")
	// 	return nil, err
	// }

	_, err = s.inventoryGateway.DeductQuantity(context.Background(), createdOrder)
	if err != nil {
		return nil, err
	}

	return createdOrder, nil
}

func (s *OrderService) UpdateOrder(dto dto.UpdateOrderRequestDto) (*models.Order, error) {
	order, err := s.orderRepository.GetByID(dto.ID)
	if err != nil {
		return nil, err
	}
	order.Status = string(dto.Status)
	return order, s.orderRepository.Update(order)
}

func (s *OrderService) PaymentOrderCompleted(dto dto.PaymentResponse) (*models.Order, error) {

	order, err := s.orderRepository.GetByID(dto.OrderID)
	if err != nil {
		return nil, err
	}
	switch dto.Status {
	case "COMPLETED":
		order.Status = "PAID"
	case "FAILED":
		order.Status = "PAYMENT_FAILED"
		_, err = s.inventoryGateway.RefundQuantity(context.Background(), order)
		if err != nil {
			return nil, err
		}
	case "PENDING":
		order.Status = "AWAITING_PAYMENT"
	default:
		order.Status = "UNKNOWN"
	}
	order.PaymentID = dto.ID
	return order, s.orderRepository.Update(order)
}
